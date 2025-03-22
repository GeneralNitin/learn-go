package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Starting the application...")
	r.HandleFunc("/api/v1/avenger/{id}", getAvenger).Methods("GET")
	http.ListenAndServe(":8080", r)
	fmt.Println("Application started successfully!")
	fmt.Println("Listening on port 8080...")
}
