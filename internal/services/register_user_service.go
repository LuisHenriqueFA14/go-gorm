package services

import (
	"errors"
	"strings"
	"encoding/json"

	db "github.com/LuisHenriqueFA14/go-gorm/internal/database"
	"github.com/LuisHenriqueFA14/go-gorm/internal/models"

	"github.com/google/uuid"
)

type user struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RegisterUserService struct {}

func (s RegisterUserService) Execute(name, email, password string) ([]byte, error) {
	if len(strings.Split(email, "@")) != 2 || len(strings.Split(strings.Split(email, "@")[1], "."))  != 2 {
		return nil, errors.New("Invalid email")
	}

	if len(password) < 6 {
		return nil, errors.New("Password must have at least 6 characters")
	}

	userAlreadyExists := db.Db.First(&models.User{}, "email = ?", email)

	if userAlreadyExists.Error == nil {
		return nil, errors.New("User already exists")
	}

	u := models.User {
		Id: uuid.NewString(),
		Name: name,
		Email: email,
		Password: password,
	}

	db.Db.Create(&u)

	res, err := json.Marshal(user {
		Id: u.Id,
		Name: u.Name,
		Email: u.Email,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
