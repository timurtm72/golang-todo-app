package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/timurtm72/todo-app"
	"github.com/timurtm72/todo-app/pkg/handler"
	"github.com/timurtm72/todo-app/pkg/repository"
	"github.com/timurtm72/todo-app/pkg/service"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed to init config: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to init db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
