package http_server

import (
	"github.com/jsagl/go-from-scratch/storage"
	"net/http"
)

func NewHTTPServer(store storage.ArticleStore) {
	articleHandler := NewArticleHandler(store)

	http.HandleFunc("/", articleHandler.GetArticle)
	http.ListenAndServe(":8080", nil)
}
