package storage

import "github.com/jsagl/go-from-scratch/models"

type RecipeStore interface {
	GetById(id int64) (*models.Recipe, error)
}