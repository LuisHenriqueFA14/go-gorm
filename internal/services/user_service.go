package services

import (
	"encoding/json"
	"strings"

	db "github.com/LuisHenriqueFA14/go-gorm/internal/database"
)

type user struct {
	Name        string `json:"name"`
	Description	string `json:"description"`
}

type UserService struct {}

func (s *UserService) Execute(user_name string) ([]byte, error) {
	u := user{}
	
	db.Db.Where("name = ?", strings.ToLower(user_name)).First(&u)

	return json.Marshal(user {
		Name: u.Name,
		Description: u.Description,
	})
}
