package handlers

import (
	"fmt"
	"net/http"
)

func GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "All users endpoint hit")
}

func GetOneProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "One user endpoint hit")
}

func UpdateProfileDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Update User Endpoint")
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	// At some point the accounts and profiles tables need to be linked,
	// in hindsight they should never have been split.
	// However looking forward if querying this table leads to err
	// make a redirect to this func, wooooo!!!
	fmt.Fprint(w, "Create User Endpoint hit")
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete user endpoint hit")
}
