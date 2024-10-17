package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// set up for a new router and defining an endpoint for the post request
	router := mux.NewRouter()
	router.HandleFunc("/find-pairs", findPairs).Methods("POST")

	http.ListenAndServe(":8080", router)
}
