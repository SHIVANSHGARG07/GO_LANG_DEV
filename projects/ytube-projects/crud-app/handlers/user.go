package handlers

import (
	"encoding/json"
	"github.com/SHIVANSHGARG07/crud-app/models"
	"github.com/gorilla/mux"
	"net/http"
)

// in memory slices in go
var users []models.User

// r.Method      // "GET", "POST", "PUT", "DELETE"
// r.URL         // "/users/123"
// r.Body        // Request body (JSON data)
// r.Header      // Request headers
func GetAllUsers(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-Type", "application/json")

	// NewEncoder: Creates en encoder that will write to resp

	// Encode: Takes the users and convert to json
	json.NewEncoder(resp).Encode(users)

}

func GetSingleUser(resp http.ResponseWriter, req *http.Request) {

	// set response header
	resp.Header().Set("Content-Type", "application/json")

	// fetch params, could be multiple
	params := mux.Vars(req)
	id := params["id"]

	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(resp).Encode(user)
			return
		}
	}

	resp.WriteHeader(http.StatusNotFound)
	json.NewEncoder(resp).Encode(map[string]string{"error": "User not found"})

}

// decode mtlb json -> go
// Decode takes json and convert to go struct

func CreateUser(resp http.ResponseWriter, req *http.Request) {

	// set response header
	resp.Header().Set("Content-Type", "application/json")

	var user models.User

	// decode json from request body
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	// after decoding
	// user.Name  // "Shivansh"
	// user.Email // "test@gmail.com"

	// add to users slice
	users = append(users, user)

	resp.WriteHeader(http.StatusCreated)
	json.NewEncoder(resp).Encode(user)

}

// UpdateUser updates an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	// Find and update user
	for i, user := range users {
		if user.ID == id {
			// Decode new data
			var updatedUser models.User
			err := json.NewDecoder(r.Body).Decode(&updatedUser)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
				return
			}

			// Keep same ID, update other fields
			updatedUser.ID = id
			users[i] = updatedUser

			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	// User not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
}
