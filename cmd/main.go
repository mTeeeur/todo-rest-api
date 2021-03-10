package main

import (
	"log"
	todo_rest_api "todo-rest-api"
	"todo-rest-api/pkg/handler"
)

func main() {
	server := new(todo_rest_api.Server)
	handlers := new(handler.Handler)

	if err := server.Run("8000", handlers.Init()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
