package main

import (
	"github.com/jsagl/go-from-scratch/http_server"
	"github.com/jsagl/go-from-scratch/storage"
	"github.com/jsagl/go-from-scratch/usecase"
	"go.uber.org/zap"
)

func main() {
	l, _ := zap.NewDevelopment()
	defer l.Sync()
	logger := l.Sugar()

	logger.Infow("Starting application...")

	logger.Infow("Opening connection to postgres database...")
	connection, err := storage.NewPostgresConnection()
	if err != nil {
		logger.Panicw("Connection to postgres database failed", "error", err,)
	}


	logger.Infow("Postgres recipe store initializing...")
	recipeStore := storage.NewPostgresRecipeStore(connection)

	logger.Infow("Setting Recipe usecase...")
	recipeUseCase := usecase.NewRecipeUseCase(recipeStore)

	http_server.NewHTTPServer(logger, recipeUseCase)
}
