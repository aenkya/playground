package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"enkya.org/playground/cmd/todo/db"
	"enkya.org/playground/cmd/todo/model"
	"enkya.org/playground/cmd/todo/service"
)

func main() {
	log.Print("start the todo app")

	ts := service.NewService(db.New())
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")

		switch strings.ToLower(cmd[0]) {
		case "add":
			todo, err := model.NewToDo(strings.Join(cmd[1:], " "))
			if err != nil {
				fmt.Printf("error adding todo: %v\n", err.Error())
				continue
			}

			_, err = ts.Add(todo)
			if err != nil {
				fmt.Println("was unable to add this todo item")
			}
		case "delete":
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("Please provide a valid ID")
				continue
			}

			if err := ts.Delete(id); err != nil {
				fmt.Printf("item with id %d not found", id)
			}
		case "update":
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("Please provide a valid ID")
				continue
			}

			_, err = ts.Update(&model.ToDo{
				ID:    id,
				Label: strings.Join(cmd[2:], " "),
			})
			if err != nil {
				fmt.Printf("failed to update the todo item, %v", err.Error())
			}
		case "get":
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("Please provide a valid ID")
				continue
			}

			todo := ts.GetByID(id)
			if todo == nil {
				fmt.Printf("todo item with ID %d not found", id)
				continue
			}

			fmt.Println(todo.String())
		case "getall":
			todos := ts.GetAll()
			if len(todos) == 0 {
				fmt.Println("No todos available")
				continue
			}

			for _, todo := range todos {
				fmt.Println(todo.String())
			}
		default:
			fmt.Printf("invalid operation %v\n", cmd[0])
		}
	}
}
