package main

import (
	"log"
	todo_rest_api "todo-rest-api"
)

func main() {
	server := new(todo_rest_api.Server)

	if err := server.Run("8000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
