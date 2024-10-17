package main

import (
	"encoding/json"
	"net/http"
)

// controller function for find-pairs endpoint
func findPairs(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid request body. Ensure 'numbers' is an array of integers and 'target' is an integer."})
		return
	}

	// Validate that numbers is not empty
	if len(requestBody.Numbers) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseBody{Solutions: [][]int{}})
		return
	}

	// Call the function to get the index pairs
	solutions := twoSum(requestBody.Numbers, requestBody.Target)

	// Return the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseBody{Solutions: solutions})
}
