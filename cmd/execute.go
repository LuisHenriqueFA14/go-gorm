package cmd

import (
	"fmt"
	"net/http"
	"github.com/LuisHenriqueFA14/go-gorm/internal/controllers"
)

func Execute() {
	registerUserController := controllers.RegisterUserController{}

	http.HandleFunc("/users/register", registerUserController.Handle)

	fmt.Println("🚀 Server is running!")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}
