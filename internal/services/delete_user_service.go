package services

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/LuisHenriqueFA14/go-gorm/internal/auth"
	db "github.com/LuisHenriqueFA14/go-gorm/internal/database"
	"github.com/LuisHenriqueFA14/go-gorm/internal/models"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt"
)

type delete_response struct {
	Message  string `json:"message"`
}

type DeleteUserService struct {}

func (s DeleteUserService) Execute(token, password string) ([]byte, error) {
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Invalid password")
	}

	db.Db.Unscoped().Delete(&user)

	response, err := json.Marshal(delete_response{Message: "User deleted"})

	if err != nil {
		return nil, errors.New("Couldn't marshal response")
	}

	return response, nil
}
