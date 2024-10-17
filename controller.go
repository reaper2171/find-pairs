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
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid request body found. Please pass valid 'numbers' and 'target' input."})
		return
	}

	// Handle empty array edge case
	if len(requestBody.Numbers) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseBody{Solutions: [][]int{}})
		return
	}

	solutions := TwoSum(requestBody.Numbers, requestBody.Target)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseBody{Solutions: solutions})
}
