package helper

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
