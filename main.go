package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Set up routes
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/welcome", VerifyToken).Methods("GET")

	// Start server
	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
