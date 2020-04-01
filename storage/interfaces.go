package storage

import "github.com/jsagl/go-from-scratch/models"

type ArticleStore interface {
	GetArticle() (*models.Article, error)
}