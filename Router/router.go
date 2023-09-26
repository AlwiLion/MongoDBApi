package router

import (
	controller "main/Controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovie).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeletAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeletALLMovie).Methods("DELETE")
	return router
}
