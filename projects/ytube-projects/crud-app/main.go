package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SHIVANSHGARG07/crud-app/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// create new router
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetSingleUser).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")

	// start server
	fmt.Println("Server running on port 7777....")

	// print + exit
	log.Fatal(http.ListenAndServe(":7777", router))

}
