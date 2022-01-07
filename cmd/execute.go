package cmd

import (
	"fmt"
	"net/http"

	"github.com/LuisHenriqueFA14/go-gorm/internal/database"
	"github.com/LuisHenriqueFA14/go-gorm/internal/controllers"
	"github.com/LuisHenriqueFA14/go-gorm/internal/auth"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func Execute() {
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.CreateDatabase(db)

	auth.LoadJWTSecret()

	registerUserController := controllers.RegisterUserController{}
	deleteUserController := controllers.DeleteUserController{}
	listUsersController := controllers.ListUsersController{}
	userController := controllers.UserController{}
	loginUserController := controllers.LoginUserController{}
	updateUserController := controllers.UpdateUserController{}

	http.HandleFunc("/users/register", registerUserController.Handle)
	http.HandleFunc("/users/delete", deleteUserController.Handle)
	http.HandleFunc("/users/list", listUsersController.Handle)
	http.HandleFunc("/users/login", loginUserController.Handle)
	http.HandleFunc("/user/", userController.Handle)
	http.HandleFunc("/users/update", updateUserController.Handle)

	fmt.Println("🚀 Server is running!")
	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}
