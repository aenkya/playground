//go:build !codeanalysis
// +build !codeanalysis

package syntax

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDriverPractice() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	} else {
		log.Println("Environment variables loaded successfully")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		fmt.Print("MONGODB_URI not found")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Print(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			fmt.Print(err)
		}
	}()

	coll := client.Database("sample_mflix").Collection("movies")
	title := "The Dark Knight"

	var result Movie
	if err = coll.FindOne(context.TODO(), bson.M{"title": title}).Decode(&result); err != nil {
		fmt.Println("No movie found")
	}

	// add aggregate query
	pipeline := []bson.M{
		{"$match": bson.M{"title": title}},
		{"$project": bson.M{"title": 1, "year": 1, "genres": 1}},
	}
	opts := options.Aggregate().SetMaxTime(2 * time.Second)

	cursor, err := coll.Aggregate(context.Background(), pipeline, opts)
	if err != nil {
		fmt.Print(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var result bson.M
		if err = cursor.Decode(&result); err != nil {
			fmt.Print(err)
		}

		fmt.Printf("%v\n", result)
	}

	if err = cursor.Err(); err != nil {
		fmt.Print(err)
	}

	// add update query
	filter := bson.M{"title": title}
	update := bson.M{"$set": bson.M{"title": "The Dark Knight Rises"}}

	updateResult, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// add insert query
	newMovie := Movie{
		ID:           "",
		Title:        "The Dark Knight",
		Genres:       []string{"Action", "Crime", "Drama"},
		Runtime:      152,
		Year:         2008,
		Released:     time.Date(2008, time.July, 18, 0, 0, 0, 0, time.UTC),
		DateModified: time.Time{},
	}

	insertResult, err := coll.InsertOne(context.Background(), newMovie)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(insertResult.InsertedID)

	jsonData, err := json.MarshalIndent(result, "", "	")
	if err != nil {
		fmt.Print(err)
	}

	var mutex = &sync.Mutex{}
	go func() {
		mutex.Lock()
		defer mutex.Unlock()

		result.Title = "The Duck Knight Rises"

		jsonData, err = json.MarshalIndent(result, "", "	")
		if err != nil {
			fmt.Print(err)
		}

		fmt.Printf("%s\n", jsonData)
	}()

	go func() {
		mutex.Lock()
		defer mutex.Unlock()

		result.Title = "The Dark Knight Rizesss"

		jsonData, err = json.MarshalIndent(result, "", "	")
		if err != nil {
			fmt.Print(err)
		}

		fmt.Printf("%s\n", jsonData)
	}()

	http.HandleFunc("/movies", GetMovies)
	fmt.Println("Server started")

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Print(err)
	}
}

//nolint:govet // ignore struct field order
type Movie struct {
	ID           string    `bson:"_id,omitempty"`
	Title        string    `bson:"title"`
	Genres       []string  `bson:"genres,omitempty"`
	Runtime      int       `bson:"runtime,omitempty"`
	Year         int       `bson:"year,omitempty"`
	Released     time.Time `bson:"released,omitempty"`
	DateModified time.Time `bson:"dateModified,required"`
}

func GetMovies(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(`{"message": "get movies"}`)); err != nil {
		fmt.Println(err)
	}
}
