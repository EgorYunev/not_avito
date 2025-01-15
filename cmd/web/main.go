package main

import (
	"flag"
	"log"

	"github.com/EgorYunev/not_avito/config"
	"github.com/EgorYunev/not_avito/internal/data"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Application struct {
	Server *mux.Router
	Logger *zap.Logger
}

func main() {

	config.ServerPort = *flag.String("port", ":8080", "Server port")

	app := &Application{}
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	app.Logger = logger
	logger.Info("Connecting to database")
	data.Start()

	logger.Info("Starting server")

	log.Fatal(app.start())
}
