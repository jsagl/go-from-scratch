package http_server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jsagl/go-from-scratch/usecase"
	"net/http"
	"strconv"
)

type RecipeHandler struct {
	usecase usecase.RecipeUseCaseInterface
}

func NewRecipeHandler(usecase usecase.RecipeUseCaseInterface) *RecipeHandler {
	return &RecipeHandler{usecase: usecase}
}

func (handler *RecipeHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	recipe, err := handler.usecase.GetById(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(recipe)
}

func (handler *RecipeHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	recipes, err := handler.usecase.FindAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(recipes[0])
}