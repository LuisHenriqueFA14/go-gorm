package controllers

import (
	"net/http"
	"strconv"

	"github.com/LuisHenriqueFA14/go-gorm/internal/services"
)

type ListUsersController struct {}

func (c ListUsersController) Handle(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		limitInt = 10
	}

	listUsersService := services.ListUsersService{}

	users, err := listUsersService.Execute(limitInt)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)
}
