package main

import (
	"log"
	todo_rest_api "todo-rest-api"
	"todo-rest-api/pkg/handler"
	"todo-rest-api/pkg/repository"
	"todo-rest-api/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(todo_rest_api.Server)

	if err := server.Run("8000", handlers.Init()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
