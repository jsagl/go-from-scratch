package storage

import (
	"database/sql"
	"fmt"
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

func (store *PostgresRecipeStore) FindAll() ([]*models.Recipe, error) {
	query:= "SELECT id, title, description, ingredients_list FROM recipes"
	rows, err := store.Connection.Query(query)
	if err != nil {
		return nil, models.ErrInternalServerError
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	recipes := make([]*models.Recipe, 0)

	var ingredients string

	for rows.Next() {
		recipe := &models.Recipe{}
		err := rows.Scan(
			&recipe.ID,
			&recipe.Title,
			&recipe.Description,
			&ingredients,
		)

		recipe.IngredientsList = strings.Split(ingredients, ";")

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}