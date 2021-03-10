package main

import (
	"github.com/spf13/viper"
	"log"
	todo_rest_api "todo-rest-api"
	"todo-rest-api/pkg/handler"
	"todo-rest-api/pkg/repository"
	"todo-rest-api/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error while readfing config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(todo_rest_api.Server)

	if err := server.Run(viper.GetString("port"), handlers.Init()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
