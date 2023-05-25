//go:build !codeanalysis
// +build !codeanalysis

package golang

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type myType int

type MyReader struct{}

func (r *MyReader) Read(p []byte) (int, error) {
	// Fill the byte slice with 'A' until it is full
	for i := 0; ; i++ {
		p[i%len(p)] = 'A'
		return len(p), nil
	}
}

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	for k, v := range p[:n] {
		p[k] = rot13(v)
	}

	return n, err
}

//nolint:gocritic // the if statement makes more sense here
func rot13(c byte) byte {
	const rot = 13
	if c >= 'a' && c <= 'z' {
		if c+rot <= 'z' {
			return c + rot
		}

		return c - rot
	} else if c >= 'A' && c <= 'Z' {
		if c+rot <= 'Z' {
			return c + rot
		}

		return c - rot
	} else {
		return c
	}
}

func GoPractice() {
	i := 8

	var x myType = 8
	if i%2 == 0 {
		fmt.Printf("Even: %T\n", x)
	} else {
		fmt.Println("Odd")
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Hello from another goroutine")
	}()

	fmt.Println(runtime.NumCPU())

	arr := [...]int{3, 5, 2}
	y := arr[0]
	fmt.Println(x, len(arr), y)

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println(sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x", sample[i]) // hexa decimal representation
	}

	fmt.Println()

	for _, rune := range sample {
		fmt.Printf("% d\n", rune)
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	if _, err := io.Copy(os.Stdout, &r); err != nil {
		fmt.Print(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
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

	pings := make(chan string)
	pongs := make(chan string)

	go pinger(pings)
	go ponger(pings, pongs)
	go printer(pongs)

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Print(err)
	}
}

func pinger(pings chan<- string) {
	for i := 0; i < 5; i++ {
		pings <- "ping"
	}
}

func ponger(pings <-chan string, pongs chan<- string) {
	for {
		select {
		case msg := <-pings:
			pongs <- msg
		default:
			pongs <- "pong"
		}
	}
}

func printer(pongs <-chan string) {
	for {
		pong := <-pongs
		fmt.Println(pong)
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
