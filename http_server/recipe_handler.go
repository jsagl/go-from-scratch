package http_server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jsagl/go-from-scratch/models"
	"github.com/jsagl/go-from-scratch/usecase"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type RecipeHandler struct {
	logger *zap.SugaredLogger
	usecase usecase.RecipeUseCaseInterface
}

func NewRecipeHandler(logger *zap.SugaredLogger, usecase usecase.RecipeUseCaseInterface) *RecipeHandler {
	return &RecipeHandler{usecase: usecase, logger: logger}
}

func (handler *RecipeHandler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	handler.logger.Infow("HTTP request", "method", r.Method,"route", r.URL)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrInvalidIdInURL)
		//fmt.Println(models.ErrInvalidIdInURL)
		return
	}

	recipe, err := handler.usecase.GetById(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipe)
}

func (handler *RecipeHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	handler.logger.Infow("HTTP request", "method", r.Method,"route", r.URL)
	recipes, err := handler.usecase.FindAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(recipes[0])
}

