package cmd

import (
	"fmt"
	"net/http"
)

func Execute() {
	fmt.Println("🚀 Server is running!")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}
