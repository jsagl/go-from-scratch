package usecase

import (
	"github.com/jsagl/go-from-scratch/models"
	"github.com/jsagl/go-from-scratch/storage"
)

type RecipeUseCaseInterface interface {
	GetById(id int64) (*models.Recipe, error)
	FindAll() ([]*models.Recipe, error)
}

type RecipeUseCase struct  {
	store storage.RecipeStore
}

func NewRecipeUseCase(store storage.RecipeStore) RecipeUseCaseInterface {
	return &RecipeUseCase{store: store}
}

func (usecase *RecipeUseCase) GetById(id int64) (*models.Recipe, error) {
	recipe, err := usecase.store.GetById(id)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (usecase *RecipeUseCase) FindAll() ([]*models.Recipe, error) {
	recipes, err := usecase.store.FindAll()
	if err != nil {
		return nil, err
	}

	return recipes, nil
}