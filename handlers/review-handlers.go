package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jpartridge95/go-app-v1/database"

	"github.com/jpartridge95/go-app-v1/model"

	_ "github.com/go-sql-driver/mysql"
)

// Unnecessary endpoint here
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "you have reached the homepage")
}

// Will be used on a "review feed"
func AllReviewsSummary(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	var ResultsSlice []model.ReviewSummary

	results, err := db.Query("SELECT productName, picture, score, personID FROM reviews")
	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var query model.ReviewSummary
		err := results.Scan(&query.ProductName, &query.Picture, &query.Score, &query.ProfileID)
		if err != nil {
			log.Fatal(err)
		}
		QueryEach := model.ReviewSummary{
			ProductName: query.ProductName,
			Picture:     query.Picture,
			Score:       query.Score,
			ProfileID:   query.ProfileID,
		}
		ResultsSlice = append(ResultsSlice, QueryEach)
	}

	json.NewEncoder(w).Encode(ResultsSlice)

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
