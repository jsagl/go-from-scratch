package http_server

import (
	"github.com/jsagl/go-from-scratch/usecase"
	"net/http"
)

func NewHTTPServer(usecase usecase.ArticleUseCaseInterface) {
	articleHandler := NewArticleHandler(usecase)

	http.HandleFunc("/", articleHandler.GetArticle)
	http.ListenAndServe(":8080", nil)
}
