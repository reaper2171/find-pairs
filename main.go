package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reaper2171/find-pairs/controllers"
)

func main() {
	// set up for a new router and defining an endpoint for the post request
	router := mux.NewRouter()
	router.HandleFunc("/find-pairs", controllers.FindPairs).Methods("POST")

	log.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
