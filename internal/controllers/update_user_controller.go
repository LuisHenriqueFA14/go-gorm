package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/LuisHenriqueFA14/go-gorm/internal/services"
)

type update_user struct {
	Description string  `json:"description"`
	Password    string  `json:"password"`
}

type UpdateUserController struct {}

func (c UpdateUserController) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	u := update_user{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if u.Description == "" && u.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing fields"))
		return
	}

	updateUserService := services.UpdateUserService{}

	token := strings.Split(r.Header.Get("Authorization"), " ")[1]

	response, err := updateUserService.Execute(token, u.Description, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
