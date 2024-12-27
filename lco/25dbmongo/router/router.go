package router

import (
	"github.com/GeneralNitin/mongodb/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteAllMovie).Methods("DELETE")

	return router
}
