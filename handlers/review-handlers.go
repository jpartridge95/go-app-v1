package handlers

import (
	"fmt"
	"net/http"
)

// Unnecessary endpoint here
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "you have reached the homepage")
}

// Will be used on a "review feed"
func AllReviews(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "All Reviews Endpoint")
}

// Create a review search

// will be used to generate a full review, map etc.
func OneReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "One Review Endpoint")
}

// Self explanatory
func PostReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post Review Endpoint")
}

// Self explanatory, Only available to the post's author
func EditReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit Review Endpoint Hit")
}

// Self explanatory again, must be only deletable by either a moderator or the author
func DeleteReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Review Endpoint Hit")
}
