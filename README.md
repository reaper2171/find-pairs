# Two-Sum Backend API

This is a backend API built in Go that solves the two-sum problem. 
It provides an endpoint to find all pairs of indices in an array whose values sum up to a given target.

# Features

- Find all pairs of indices in an array of integers that sum to a target value.
- Handles multiple requests concurrently.
- Validates input data for correctness.

## Endpoints

### `POST /find-pairs`

This endpoint receives a JSON object with an array of integers and a target sum.


## Installation
- Clone the repostiory
- Change the directory inside the project and run "go mod tidy"
- Run "go run main.go"