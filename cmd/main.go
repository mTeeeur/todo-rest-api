package main

import (
	_ "github.com/lib/pq"
	todo_rest_api "github.com/mTeeeur/todo-rest-api"
	"github.com/mTeeeur/todo-rest-api/pkg/handler"
	"github.com/mTeeeur/todo-rest-api/pkg/repository"
	"github.com/mTeeeur/todo-rest-api/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error while readfing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "admin",
	})

	if err != nil {
		log.Fatalf("Error while opening database connection: %s", err.Error())
	}

	repos := repository.NewRepository(db)
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
