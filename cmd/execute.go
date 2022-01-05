package cmd

import (
	"fmt"
	"net/http"

	"github.com/LuisHenriqueFA14/go-gorm/internal/database"
	"github.com/LuisHenriqueFA14/go-gorm/internal/controllers"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func Execute() {
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.CreateDatabase(db)

	registerUserController := controllers.RegisterUserController{}
	deleteUserController := controllers.DeleteUserController{}
	listUsersController := controllers.ListUsersController{}

	http.HandleFunc("/users/register", registerUserController.Handle)
	http.HandleFunc("/users/delete", deleteUserController.Handle)
	http.HandleFunc("/users/list", listUsersController.Handle)

	fmt.Println("🚀 Server is running!")
	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}
