package http_server

import (
	"github.com/gorilla/mux"
	"github.com/jsagl/go-from-scratch/usecase"
	"net/http"
)

func NewHTTPServer(usecase usecase.RecipeUseCaseInterface) {
	router := mux.NewRouter()

	recipeHandler := NewRecipeHandler(usecase)

	router.HandleFunc("/recipes/{id}", recipeHandler.GetById)
	http.ListenAndServe(":8080", router)
}
