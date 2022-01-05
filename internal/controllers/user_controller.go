package controllers

import (
	"net/http"
	"strings"

	"github.com/LuisHenriqueFA14/go-gorm/internal/services"
)

type UserController struct {}


func (c UserController) Handle(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/")

	if url[2] == "" {
		w.Write([]byte("User not found"))
		return;
	}

	userService := services.UserService{}

	response, err := userService.Execute(url[2])

	if err != nil {
		w.Write([]byte(err.Error()))
		return;
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
