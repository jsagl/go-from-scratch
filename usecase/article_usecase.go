package usecase

import (
	"github.com/jsagl/go-from-scratch/models"
	"github.com/jsagl/go-from-scratch/storage"
)

type ArticleUseCaseInterface interface {
	GetArticle() *models.Article
}

type ArticleUseCase struct  {
	store storage.ArticleStore
}

func NewArticleUseCase(store storage.ArticleStore) ArticleUseCaseInterface {
	return &ArticleUseCase{store: store}
}

func (usecase *ArticleUseCase) GetArticle() *models.Article {
	article, _ := usecase.store.GetArticle()
	return article
}