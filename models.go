package main

// defined the request body schema here
type RequestBody struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

// defined the respons body schema here
type ResponseBody struct {
	Solutions [][]int `json:"solutions"`
}

// defined the error response body schema here
type ErrorResponse struct {
	Message string `json:"message"`
}
