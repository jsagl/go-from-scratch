package http_server

import (
	"fmt"
	"github.com/jsagl/go-from-scratch/storage"
	"net/http"
)

type ArticleHandler struct {
	store storage.ArticleStore
}

func NewArticleHandler(store storage.ArticleStore) *ArticleHandler {
	return &ArticleHandler{store: store}
}

func (handler *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	article, _ := handler.store.GetArticle()

	fmt.Println(article.Title)
}