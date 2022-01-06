package services

import (
	"encoding/json"
	"errors"
	"time"

	db "github.com/LuisHenriqueFA14/go-gorm/internal/database"
	"github.com/LuisHenriqueFA14/go-gorm/internal/models"
	"github.com/LuisHenriqueFA14/go-gorm/internal/auth"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type login_response struct {
	Token  string `json:"token"`
}

type LoginUserService struct {}

type JwtClaim struct {
	Id string
	jwt.StandardClaims
}

func (s LoginUserService) Execute(name, email, password string) ([]byte, error) {
	var user models.User

	if email != "" {
		db.Db.Where("email = ?", email).First(&user)
	} else if name != "" {
		db.Db.Where("name = ?", name).First(&user)
	} else {
		return nil, errors.New("Invalid parameters")
	}

	if user.Id == "" {
		return nil, errors.New("User not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.New("Wrong password")
	}

	claims := &JwtClaim{
		Id: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(auth.JWTToken))

	return json.Marshal(login_response{
		Token: tokenString,
	})
}
