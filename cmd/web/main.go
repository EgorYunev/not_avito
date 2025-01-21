package main

import (
	"flag"
	"log"

	"github.com/EgorYunev/not_avito/config"
	"github.com/EgorYunev/not_avito/internal/data"
	"github.com/EgorYunev/not_avito/internal/services"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Application struct {
	Server      *mux.Router
	Logger      *zap.Logger
	UserService *services.UserService
	AdService   *services.AdService
}

func main() {

	config.ServerPort = *flag.String("port", ":8080", "Server port")

	app := &Application{
		UserService: &services.UserService{
			UserRepository: &data.UserRepository{},
		},
		AdService: &services.AdService{
			Repo: &data.AdRepository{},
		},
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	app.Logger = logger
	logger.Info("Connecting to database")
	db := data.Start()
	defer db.Close()

	app.UserService.UserRepository.DB = db
	app.AdService.Repo.DB = db

	logger.Info("Starting server")

	log.Fatal(app.start())
}
