package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "you have reached the homepage")
}

func AllReviews(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "All Reviews Endpoint")
}

func OneReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "One Review Endpoint")
}

func PostReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post Review Endpoint")
}

func EditReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit Review Endpoint Hit")
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Review Endpoint Hit")
}
