package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LuisHenriqueFA14/go-gorm/internal/services"
)

type login_user struct {
	Name string     `json:"name"`
	Email string    `json:"email"`
	Password string `json:"password"`
}

type LoginUserController struct {}


func (c LoginUserController) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	loginUserService := services.LoginUserService{}

	u := login_user{}
	
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if (u.Name == "" && u.Email == "") || u.Password == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}

	response, err := loginUserService.Execute(u.Name, u.Email, u.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
