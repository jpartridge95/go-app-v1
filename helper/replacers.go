package helper // Should make a new helper to set response headers

import "github.com/jpartridge95/go-app-v1/model"

func Replacer(old model.AccountChange, new *model.AccountChange) model.AccountChange {

	firstName := &new.FirstName
	lastname := &new.LastName
	email := &new.Email
	phone := &new.PhoneNumber
	dob := &new.DateOfBirth
	secQ := &new.SecurityQuestion

	if new.FirstName == "" {
		*firstName = old.FirstName
	}

	if new.LastName == "" {
		*lastname = old.LastName
	}

	if new.Email == "" {
		*email = old.Email
	}

	if new.PhoneNumber == "" {
		*phone = old.PhoneNumber
	}

	if new.DateOfBirth == "" {
		*dob = old.DateOfBirth
	}

	if new.SecurityQuestion == "" {
		*secQ = old.SecurityQuestion
	}

	return *new
}

func ProfileReplacer(old model.Profile, new *model.Profile) model.Profile {
	username := &new.UserName
	age := &new.Age
	city := &new.City
	accountID := old.Accountid

	if new.UserName == "" {
		*username = old.UserName
	}
	if new.Age == 0 {
		*age = old.Age
	}
	if new.City == "" {
		*city = old.City
	}

	return *new
}

// potential for an iknterface, If I find myself making more of these methods
// I will be creating an inteface for these methods
