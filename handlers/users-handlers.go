package handlers

import (
	"fmt"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "All users endpoint hit")
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "One user endpoint hit")
}

func UpdateUserDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Update User Endpoint")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create User Endpoint hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete user endpoint hit")
}
