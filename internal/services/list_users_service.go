package services

import (
	"encoding/json"

	db "github.com/LuisHenriqueFA14/go-gorm/internal/database"
)

type list_users struct {
	Name string `json:"name"`
}

type ListUsersService struct {}

func (s ListUsersService) Execute(limit int) ([]byte, error) {
	type user list_users

	var users []user

	db.Db.Limit(limit).Find(&users)

	return json.Marshal(users)
}
