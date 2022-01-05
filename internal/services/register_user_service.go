package services

import (
	"errors"
	"strings"
	"encoding/json"

	db "github.com/LuisHenriqueFA14/go-gorm/internal/database"
	"github.com/LuisHenriqueFA14/go-gorm/internal/models"

	"github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
)

type register_user struct {
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

	if strings.Contains(name, " ") {
		return nil, errors.New("Name cannot contain spaces")
	}

	userAlreadyExistsEmail := db.Db.First(&models.User{}, "email = ?", email)
	userAlreadyExistsName := db.Db.First(&models.User{}, "name = ?", strings.ToLower(name))

	if userAlreadyExistsEmail.Error == nil {
		return nil, errors.New("Email already in use")
	}

	if userAlreadyExistsName.Error == nil {
		return nil, errors.New("Name already in use")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	hashPassword := string(hash)

	u := models.User {
		Id: uuid.NewString(),
		Name: strings.ToLower(name),
		Email: email,
		Password: hashPassword,
	}

	db.Db.Create(&u)

	res, err := json.Marshal(register_user {
		Id: u.Id,
		Name: u.Name,
		Email: u.Email,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
