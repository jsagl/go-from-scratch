package main

import (
	"github.com/jsagl/go-from-scratch/http_server"
	"github.com/jsagl/go-from-scratch/storage"
	"github.com/jsagl/go-from-scratch/usecase"
)

func main() {
	connection, _ := storage.NewPostgresConnection()
	recipeStore := storage.NewPostgresRecipeStore(connection)
	recipeUseCase := usecase.NewRecipeUseCase(recipeStore)
	http_server.NewHTTPServer(recipeUseCase)
}
