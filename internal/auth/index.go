package auth

import (
	"fmt"

	"github.com/joho/godotenv"

	"os"
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

	fmt.Println(JWTToken)
}
