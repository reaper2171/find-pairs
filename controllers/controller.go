package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/reaper2171/find-pairs/models"
	"github.com/reaper2171/find-pairs/utils"
)

// controller function for find-pairs endpoint
func FindPairs(w http.ResponseWriter, r *http.Request) {
	var requestBody models.RequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "Invalid request body found. Please pass valid 'numbers' and 'target' input."})
		return
	}

	// Handle empty array edge case
	if len(requestBody.Numbers) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.ResponseBody{Solutions: [][]int{}})
		return
	}

	solutions := utils.TwoSum(requestBody.Numbers, requestBody.Target)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.ResponseBody{Solutions: solutions})
}
