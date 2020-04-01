package main

import (
	"github.com/jsagl/go-from-scratch/http_server"
	"github.com/jsagl/go-from-scratch/storage"
	"github.com/jsagl/go-from-scratch/usecase"
)

func main() {
	connection, _ := storage.NewPostgresConnection()
	articleStore := storage.NewPostgresArticleStore(connection)
	articleUseCase := usecase.NewArticleUseCase(articleStore)
	http_server.NewHTTPServer(articleUseCase)
}
