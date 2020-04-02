package models

type Recipe struct {
	ID int64
	Title string
	Description string
	IngredientsList []string
}