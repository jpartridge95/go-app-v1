package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jpartridge95/go-app-v1/database"

	"github.com/jpartridge95/go-app-v1/model"

	_ "github.com/go-sql-driver/mysql"
)

// Unnecessary endpoint here to be cleaned up in a refactor
// func Preflight(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "you have reached the homepage")
// }

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
		boughtfromlat,
		boughtfromlong,
		boughtfor,
		currency,
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
		&query.BoughtFromLat,
		&query.BoughtFromLong,
		&query.BoughtFor,
		&query.Currency,
		&query.FullReview,
		&query.ProfileID,
	)
	if err != nil {
		log.Fatal(err)
	}

	QueryResult := model.Review{
		ReviewID:       query.ReviewID,
		ProductName:    query.ProductName,
		Picture:        query.Picture,
		Score:          query.Score,
		BoughtFrom:     query.BoughtFrom,
		BoughtFromLat:  query.BoughtFromLat,
		BoughtFromLong: query.BoughtFromLong,
		BoughtFor:      query.BoughtFor,
		Currency:       query.Currency,
		FullReview:     query.FullReview,
		ProfileID:      query.ProfileID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(QueryResult)
}

// Self explanatory
func PostReview(w http.ResponseWriter, r *http.Request) {

	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "multipart/form-data; boundary='boundary'")

	r.ParseMultipartForm(32 << 20)

	data := r.MultipartForm

	keyVals := data.Value

	score, err := strconv.ParseInt(keyVals["score"][0], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	latitude, err := strconv.ParseFloat(keyVals["boughtFromLat"][0], 64)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	longitude, err := strconv.ParseFloat(keyVals["boughtFromLong"][0], 64)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	price, err := strconv.ParseFloat(keyVals["pricePaid"][0], 32)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id, err := strconv.ParseInt(keyVals["createdBy"][0], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	newEntry := model.Review{
		ProductName:    keyVals["productName"][0],
		Picture:        keyVals["productImage"][0],
		Score:          score,
		BoughtFrom:     keyVals["locationBought"][0],
		BoughtFromLat:  latitude,
		BoughtFromLong: longitude,
		BoughtFor:      float32(price),
		Currency:       keyVals["currency"][0],
		FullReview:     keyVals["fullReview"][0],
		ProfileID:      id,
	}

	insert, err := db.Prepare(`INSERT INTO
		reviews (
			productName,
			picture,
			score,
			boughtfrom,
			boughtfromlat,
			boughtfromlong,
			boughtfor,
			currency,
			fullreview,
			personID
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	res, err := insert.Exec(
		newEntry.ProductName,
		newEntry.Picture,
		newEntry.Score,
		newEntry.BoughtFrom,
		newEntry.BoughtFromLat,
		newEntry.BoughtFromLong,
		newEntry.BoughtFor,
		newEntry.Currency,
		newEntry.FullReview,
		newEntry.ProfileID,
	)

	lastID, err := res.LastInsertId()

	fmt.Fprintf(w, "%d", lastID)

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
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	vars := mux.Vars(r)
	id := vars["id"]

	var result model.EditScore

	json.NewDecoder(r.Body).Decode(&result)

	editedReview := model.EditScore{
		Score: result.Score,
	}

	_, err := db.Query(
		`UPDATE
			reviews
		SET
			score = ?
		WHERE
			reviewid = ?`,
		editedReview.Score,
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, "Review score successfully changed")
}

// Self explanatory again, must be only deletable by either a moderator or the author
func DeleteReview(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Query(`
		DELETE FROM
			reviews
		WHERE
			reviewid = ?`,
		id,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, "Delete Review Endpoint Hit")
}
