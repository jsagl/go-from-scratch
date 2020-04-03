package http_server

import (
	"github.com/gorilla/mux"
	"github.com/jsagl/go-from-scratch/usecase"
	"go.uber.org/zap"
	"net/http"
)

func NewHTTPServer(logger *zap.SugaredLogger, usecase usecase.RecipeUseCaseInterface) {
	logger.Infow("Setting up router...")
	router := mux.NewRouter()

	recipeHandler := NewRecipeHandler(logger, usecase)

	router.HandleFunc("/recipes/{id}", recipeHandler.GetById).Methods("GET")
	router.HandleFunc("/recipes", recipeHandler.FindAll).Methods("GET")

	logger.Infow("HTTP Server listening on localhost:8080")
	logger.Fatal(http.ListenAndServe(":8080", router))
}
