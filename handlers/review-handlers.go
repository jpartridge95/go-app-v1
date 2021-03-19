package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	results, err := db.Query("SELECT reviewid, productName, picture, score, personID FROM reviews")
	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var query model.ReviewSummary
		err := results.Scan(&query.ReviewID, &query.ProductName, &query.Picture, &query.Score, &query.ProfileID)
		if err != nil {
			log.Fatal(err)
		}
		QueryEach := model.ReviewSummary{
			ReviewID:    query.ReviewID,
			ProductName: query.ProductName,
			Picture:     query.Picture,
			Score:       query.Score,
			ProfileID:   query.ProfileID,
		}
		ResultsSlice = append(ResultsSlice, QueryEach)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResultsSlice)
}

// Create a review search

// will be used to generate a full review, map etc.
func OneReview(w http.ResponseWriter, r *http.Request) {

	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	vars := mux.Vars(r)
	id := vars["id"]

	var query model.Review

	err := db.QueryRow(`
	SELECT
		reviewid, 
		productName, 
		picture, 
		score,
		boughtfrom,
		boughtfor,
		fullreview,
		personID
	FROM
		reviews
	WHERE 
		reviewid = ?`, id).Scan(
		&query.ReviewID,
		&query.ProductName,
		&query.Picture,
		&query.Score,
		&query.BoughtFrom,
		&query.BoughtFor,
		&query.FullReview,
		&query.ProfileID,
	)
	if err != nil {
		log.Fatal(err)
	}

	QueryResult := model.Review{
		ReviewID:    query.ReviewID,
		ProductName: query.ProductName,
		Picture:     query.Picture,
		Score:       query.Score,
		BoughtFrom:  query.BoughtFrom,
		BoughtFor:   query.BoughtFor,
		FullReview:  query.FullReview,
		ProfileID:   query.ProfileID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(QueryResult)
}

// Self explanatory
func PostReview(w http.ResponseWriter, r *http.Request) {

	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	var review model.Review

	json.NewDecoder(r.Body).Decode(&review)

	newEntry := model.Review{
		ProductName: review.ProductName,
		Picture:     review.Picture,
		Score:       review.Score,
		BoughtFrom:  review.BoughtFrom,
		BoughtFor:   review.BoughtFor,
		FullReview:  review.FullReview,
		ProfileID:   review.ProfileID,
	}

	insert, err := db.Query(`INSERT INTO 
		reviews ( 
			productName, 
			picture, 
			score,
			boughtfrom,
			boughtfor,
			fullreview,
			personID
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		newEntry.ProductName,
		newEntry.Picture,
		newEntry.Score,
		newEntry.BoughtFrom,
		newEntry.BoughtFor,
		newEntry.FullReview,
		newEntry.ProfileID,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, "New review for "+newEntry.ProductName+" created")

	insert.Close()
}

// Self explanatory, Only available to the post's author
func EditReview(w http.ResponseWriter, r *http.Request) {

	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	vars := mux.Vars(r)
	id := vars["id"]

	var result model.EditReview

	json.NewDecoder(r.Body).Decode(&result)

	editedReview := model.EditReview{
		FullReview: result.FullReview,
	}

	_, err := db.Query(
		`UPDATE
			reviews
		SET
			fullreview = ?
		WHERE
			reviewid = ?`,
		editedReview.FullReview,
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, "Review text successfully changed")
}

// Separate func to change score, also only available to post author

func EditReviewScore(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Score endpoint hit")
}

// Self explanatory again, must be only deletable by either a moderator or the author
func DeleteReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Review Endpoint Hit")
}
