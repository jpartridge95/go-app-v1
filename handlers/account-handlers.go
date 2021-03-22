package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jpartridge95/go-app-v1/database"
	"github.com/jpartridge95/go-app-v1/helper"
	"github.com/jpartridge95/go-app-v1/model"

	"github.com/gorilla/mux"
)

// Function Name explains all, gets an unfiltered summary of all accounts

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	var resultSlice []model.AccountSummary

	results, err := db.Query(`
		SELECT
			accountid,
			firstname,
			lastname,
			email,
			dob
		FROM
			accounts
	`)

	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var query model.AccountSummary
		err := results.Scan(
			&query.AccountID,
			&query.FirstName,
			&query.LastName,
			&query.Email,
			&query.DateOfBirth,
		)
		if err != nil {
			log.Fatal(err)
		}

		queryEach := model.AccountSummary{
			AccountID:   query.AccountID,
			FirstName:   query.FirstName,
			LastName:    query.LastName,
			Email:       query.Email,
			DateOfBirth: query.DateOfBirth,
		}

		resultSlice = append(resultSlice, queryEach)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultSlice)
}

// Takes a url variable ID and returns one user account associated with that ID,
// considering making it an email check?? orr implementing lower a returned id
// in create account to help with the linking of profiles to accounts.

func GetOneAccount(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	vars := mux.Vars(r)
	id := vars["id"]

	var query model.Account

	result := db.QueryRow(`SELECT * FROM accounts WHERE accountid = ?`, id)

	err := result.Scan(
		&query.AccountID,
		&query.FirstName,
		&query.LastName,
		&query.Email,
		&query.Password,
		&query.PhoneNumber,
		&query.DateOfBirth,
		&query.SecurityQuestion,
		&query.SecurityAnswer,
	)

	if err != nil {
		log.Fatal(err)
	}

	queryResult := model.Account{
		AccountID:        query.AccountID,
		FirstName:        query.FirstName,
		LastName:         query.LastName,
		Email:            query.Email,
		Password:         query.Password,
		PhoneNumber:      query.PhoneNumber,
		DateOfBirth:      query.DateOfBirth,
		SecurityQuestion: query.SecurityQuestion,
		SecurityAnswer:   query.SecurityAnswer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(queryResult)

	fmt.Fprint(w, "One user endpoint hit")
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	var account model.Account

	json.NewDecoder(r.Body).Decode(&account)

	newEntry := model.Account{
		FirstName:        account.FirstName,
		LastName:         account.LastName,
		Email:            account.Email,
		Password:         account.Password,
		PhoneNumber:      account.PhoneNumber,
		DateOfBirth:      account.DateOfBirth,
		SecurityQuestion: account.SecurityQuestion,
		SecurityAnswer:   account.SecurityAnswer,
	}

	insert, err := db.Query(`
		INSERT INTO 
			accounts (
				firstname,
				lastname,
				email,
				pass,
				phone,
				dob,
				securityquestion,
				securityanswer
			) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		newEntry.FirstName,
		newEntry.LastName,
		newEntry.Email,
		newEntry.Password,
		newEntry.PhoneNumber,
		newEntry.DateOfBirth,
		newEntry.SecurityQuestion,
		newEntry.SecurityAnswer,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "New account for "+newEntry.FirstName+" "+newEntry.LastName+" created")

	insert.Close()
}

func UpdateAccountDetails(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectionOpen()
	defer database.ConnectionClose(db)

	vars := mux.Vars(r)
	id := vars["id"]

	var account model.AccountChange
	var update model.AccountChange

	json.NewDecoder(r.Body).Decode(&update)

	result := db.QueryRow(`
		SELECT 
			firstname, 
			lastname, 
			email,
			phone,
			dob,
			securityquestion 
		FROM 
			accounts 
		WHERE 
			accountid = ?`, id,
	)
	err := result.Scan(
		&account.FirstName,
		&account.LastName,
		&account.Email,
		&account.PhoneNumber,
		&account.DateOfBirth,
		&account.SecurityQuestion,
	)

	if err != nil {
		log.Fatal(err)
	}

	accountEdit := helper.Replacer(account, &update)

	_, err = db.Query(`
		UPDATE
			accounts
		SET
			firstname = ?,
			lastname = ?,
			email = ?,
			phone = ?,
			dob = ?,
			securityquestion = ?
		WHERE
			accountid = ?`,
		accountEdit.FirstName,
		accountEdit.LastName,
		accountEdit.Email,
		accountEdit.PhoneNumber,
		accountEdit.DateOfBirth,
		accountEdit.SecurityQuestion,
		id,
	)

	fmt.Fprint(w, "Update User Endpoint")
}

func ChangePassWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have reached the change password endpoint")
}

func ChangeSecurityAnswer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the answer changer enjoy your stay")
}

func Login(w http.ResponseWriter, r *http.Request) {
	// check credentials and return profile response, then can use profile response to write reviews
	fmt.Fprintf(w, "Oh, you're trying to login are you?")
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete user endpoint hit")
}
