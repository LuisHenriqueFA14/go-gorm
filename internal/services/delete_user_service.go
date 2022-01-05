package services

import (
	"encoding/json"
	"errors"

	db "github.com/LuisHenriqueFA14/go-gorm/internal/database"
	"github.com/LuisHenriqueFA14/go-gorm/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type delete_response struct {
	Message  string `json:"message"`
}

type DeleteUserService struct {}

func (s DeleteUserService) Execute(id, email, password string) ([]byte, error) {
	if id == "" && email == "" {
		return nil, errors.New("Missing fields")
	}

	if id != "" {
		var user models.User
		db.Db.Where("id = ?", id).First(&user)

		if user.Id == "" {
			return nil, errors.New("User not found")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
			db.Db.Unscoped().Delete(&user)
			return json.Marshal(delete_response{Message: "User deleted"})
		} else {
			return nil, errors.New("Wrong password")
		}
	} else if email != "" {
		var user models.User
		db.Db.Where("email = ?", email).First(&user)

		if user.Id == "" {
			return nil, errors.New("User not found")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
			db.Db.Unscoped().Delete(&user)
			return json.Marshal(delete_response{Message: "User deleted"})
		} else {
			return nil, errors.New("Wrong password")
		}
	}

	return nil, errors.New("Invalid fields")
}
