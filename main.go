package main

import (
	"github.com/jsagl/go-from-scratch/http_server"
	"github.com/jsagl/go-from-scratch/storage"
)

func main() {
	connection, _ := storage.NewPostgresConnection()
	articleStore := storage.NewPostgresArticleStore(connection)
	http_server.NewHTTPServer(articleStore)
}
