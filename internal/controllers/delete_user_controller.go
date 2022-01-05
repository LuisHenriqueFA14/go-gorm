package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LuisHenriqueFA14/go-gorm/internal/services"
)

type delete_user struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type DeleteUserController struct {}


func (c DeleteUserController) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	deleteUserService := services.DeleteUserService{}

	u := delete_user{}

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if u.Password == "" {
		http.Error(w, "Missing password", http.StatusBadRequest)
		return
	}

	response, err := deleteUserService.Execute(u.Id, u.Email, u.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
