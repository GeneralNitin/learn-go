package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getAvenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "You've requested the avenger with ID: %s\n", id)
}
