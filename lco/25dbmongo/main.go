package main

import (
	"fmt"
	"github.com/GeneralNitin/mongodb/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started...")

	r := router.Router()

	http.ListenAndServe(":4000", r)

	fmt.Println("Listening at port 4000...")
}
