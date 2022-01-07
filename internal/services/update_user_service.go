package services

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/LuisHenriqueFA14/go-gorm/internal/auth"
	db "github.com/LuisHenriqueFA14/go-gorm/internal/database"
	"github.com/LuisHenriqueFA14/go-gorm/internal/models"

	"github.com/golang-jwt/jwt"
)

type update_response struct {
	Message string `json:"message"`
}

type UpdateUserService struct {}

func (s UpdateUserService) Execute(token string, description, password string) ([]byte, error) {
	decoded, err := jwt.ParseWithClaims(
		token,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.JWTToken), nil
		},
	)

	if err != nil {
		return nil, errors.New("Invalid Token")
	}

	claims, ok := decoded.Claims.(*JwtClaim)
	if !ok {
		return nil, errors.New("Couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	user := models.User{}
	db.Db.Where("id = ?", claims.Id).First(&user)

	if user.Id == "" {
		return nil, errors.New("User not found")
	}

	if description != "" {
		user.Description = description
	}

	if password != "" {
		user.Password = password
	}

	db.Db.Save(&user)

	response, err := json.Marshal(update_response {
		Message: "User updated",
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
