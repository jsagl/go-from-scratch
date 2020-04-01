package http_server

import (
	"fmt"
	"github.com/jsagl/go-from-scratch/usecase"
	"net/http"
)

type ArticleHandler struct {
	usecase usecase.ArticleUseCaseInterface
}

func NewArticleHandler(usecase usecase.ArticleUseCaseInterface) *ArticleHandler {
	return &ArticleHandler{usecase: usecase}
}

func (handler *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	article := handler.usecase.GetArticle()

	fmt.Println(article.Title)
}