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

	// Homepage probably going to be gone soon, we see

	r.HandleFunc("/", handlers.Home).Methods("GET")

	// These Handlers deal with Reviews and can be found in ./handlers/review-handlers.go

	r.HandleFunc("/reviews/all", handlers.AllReviewsSummary).Methods("GET")
	r.HandleFunc("/reviews/byid/{id}", handlers.OneReview).Methods("GET")
	r.HandleFunc("/reviews/create", handlers.PostReview).Methods("POST")
	r.HandleFunc("/reviews/updatereview/{id}", handlers.EditReview).Methods("PUT")
	r.HandleFunc("/reviews/updatescore/{id}", handlers.EditReviewScore).Methods("PUT")
	r.HandleFunc("/reviews/delete/{id}", handlers.DeleteReview).Methods("DELETE")

	// Handlers to deal with accounts

	r.HandleFunc("/users/all", handlers.GetAllAccounts).Methods("GET")
	r.HandleFunc("/users/byid/{id}", handlers.GetOneAccount).Methods("GET")
	r.HandleFunc("/users/create", handlers.CreateAccount).Methods("POST")
	r.HandleFunc("/users/update/{id}", handlers.UpdateAccountDetails).Methods("PUT")
	r.HandleFunc("/users/delete", handlers.DeleteAccount).Methods("DELETE")

	// Handlers to deal with Profiles

	r.HandleFunc("/profiles/all", handlers.GetAllProfiles).Methods("GET")
	r.HandleFunc("/profiles/one", handlers.GetOneProfile).Methods("GET")
	r.HandleFunc("/profiles/create", handlers.CreateProfile).Methods("POST")
	r.HandleFunc("/profiles/update", handlers.UpdateProfileDetails).Methods("PUT")
	r.HandleFunc("/profiles/delete", handlers.DeleteProfile).Methods("DELETE")

	// Start the webserver

	log.Fatal(http.ListenAndServe(":9001", r))
}
