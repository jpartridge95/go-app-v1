package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jpartridge95/go-app-v1/handlers"

	"github.com/gorilla/mux"
)

// Starts the server on port 9001 and handles all routes
// Handler functions can be found in ./handlers

func Start() {
	fmt.Printf("Application Start")
	r := mux.NewRouter().StrictSlash(true)

	// Homepage

	r.HandleFunc("/", handlers.Home).Methods("GET")

	// These Handlers deal with Reviews and can be found in ./handlers/review-handlers.go

	r.HandleFunc("/reviews/all", handlers.AllReviewsSummary).Methods("GET")
	r.HandleFunc("/reviews/byid/{id}", handlers.OneReview).Methods("GET")
	r.HandleFunc("/reviews/create", handlers.PostReview).Methods("POST")
	r.HandleFunc("/reviews/updatereview/{id}", handlers.EditReview).Methods("PUT")
	r.HandleFunc("/reviews/updatescore/{id}", handlers.EditReviewScore).Methods("PUT")
	r.HandleFunc("/reviews/delete/{id}", handlers.DeleteReview).Methods("DELETE")

	// Handlers to deal with users

	r.HandleFunc("/users/all", handlers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/one", handlers.GetOneUser).Methods("GET")
	r.HandleFunc("/users/create", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/update", handlers.UpdateUserDetails).Methods("PUT")
	r.HandleFunc("/users/delete", handlers.DeleteUser).Methods("DELETE")

	// Start the webserver

	log.Fatal(http.ListenAndServe(":9001", r))
}
