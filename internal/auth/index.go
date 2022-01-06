package auth

import (
	"os"
	
	"github.com/joho/godotenv"
)

var JWTToken string

func LoadJWTSecret() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	
	JWTToken = os.Getenv("JWT_TOKEN")

	if JWTToken == "" {
		panic("JWT_TOKEN is not defined")
	}
}
