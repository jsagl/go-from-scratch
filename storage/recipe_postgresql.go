package storage

import (
	"database/sql"
	"github.com/jsagl/go-from-scratch/models"
	"strings"
)

type PostgresRecipeStore struct {
	Connection *sql.DB
}

func NewPostgresRecipeStore(connection *sql.DB) RecipeStore {
	return &PostgresRecipeStore{Connection: connection}
}

func (store *PostgresRecipeStore) GetById(recipeId int64) (*models.Recipe, error) {
	var (
		id int64
		title string
		description string
		ingredients string
	)

	query := `SELECT id, title, description, ingredients_list FROM recipes WHERE ID = $1`
	if err := store.Connection.QueryRow(query, recipeId).Scan(&id, &title, &description, &ingredients); err != nil {
		return nil, models.ErrNotFound
	}

	ingredientsList := strings.Split(ingredients, ";")

	return &models.Recipe{ID: id, Title: title, Description: description, IngredientsList: ingredientsList}, nil
}