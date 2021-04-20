package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jpartridge95/go-app-v1/database"
	"github.com/jpartridge95/go-app-v1/helper"
	"github.com/jpartridge95/go-app-v1/model"
)

func GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	var resultSlice []model.Profile

	results, err := db.Query(`
		SELECT
			username,
			age,
			city,
			accountid
		FROM 
			persons
	`)

	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var query model.Profile
		err := results.Scan(
			&query.UserName,
			&query.Age,
			&query.City,
			&query.Accountid,
		)

		if err != nil {
			log.Fatal(err)
		}

		queryEach := model.Profile{
			UserName:  query.UserName,
			Age:       query.Age,
			City:      query.City,
			Accountid: query.Accountid,
		}

		resultSlice = append(resultSlice, queryEach)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultSlice)
}

func GetOneProfile(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	vars := mux.Vars(r)
	id := vars["id"]

	var query model.Profile

	result := db.QueryRow(`SELECT * FROM persons WHERE personID = ?`, id)

	err := result.Scan(
		&query.UserName,
		&query.Age,
		&query.City,
		&query.Accountid,
	)

	if err != nil {
		log.Fatal(err)
	}

	queryResult := model.Profile{
		UserName:  query.UserName,
		Age:       query.Age,
		City:      query.City,
		Accountid: query.Accountid,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(queryResult)
}

func UpdateProfileDetails(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	vars := mux.Vars(r)
	id := vars["id"]

	var query model.Profile
	var oldResults model.Profile

	json.NewDecoder(r.Body).Decode(&oldResults)

	results := db.QueryRow(`
		SELECT 
			*
		FROM
			persons
		WHERE
			personid = ?
	`, id)

	err := results.Scan(
		&query.UserName,
		&query.Age,
		&query.City,
		&query.Accountid,
	)

	if err != nil {
		log.Fatal(err)
	}

	profileEdit := helper.ProfileReplacer(oldResults, &query)

	_, err = db.Query(`
		UPDATE
			persons
		SET
			username = ?,
			age = ?,
			city = ?
		WHERE
			personid = ? `,
		profileEdit.UserName,
		profileEdit.Age,
		profileEdit.City,
		id,
	)

	fmt.Fprintf(w, "Profile succesfully changed")
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	// At some point the accounts and profiles tables need to be linked,
	// in hindsight they should never have been split.
	// However looking forward if querying this table for acc id leads to err
	// make a redirect to this func, wooooo!!!
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	var newProfile model.Profile

	json.NewDecoder(r.Body).Decode(&newProfile)

	_, err := db.Query(`
		INSERT INTO
			persons (
				username
				age
				city
				accountid
			)
		VALUES (?, ?, ?, ?)
	`, newProfile.UserName,
		newProfile.Age,
		newProfile.City,
		newProfile.Accountid,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, "Profile for "+newProfile.UserName+" Created")
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Query(`
		DELETE FROM 
			persons
		WHERE
			personid =?
	`, id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, "User has been deleted")
}
