package model

/*
ALL STRUCTS HERE CURRENTLY ARE USED FOR DB INTERACTION

if you came here from another file looking for a model it will be here,
bookmarks are in place to somewhat break this uninteresting file up

Contents:
	1. Profile Modelling
		(used for publicly visible info,
			username, age and city)
	2. Account modelling
		(used for any security sensitive info,
			emails, passwords, security q's etc.)
	3. Review Modelling
		(used for anything review related,
			what, where, how much, etc.)

*/

/*
--------------------------------------------------------------------
----------------------- 1. Profiles --------------------------------
--------------------------------------------------------------------
*/

type Profile struct {
	UserName  string `json:"userName"`
	Age       int32  `json:"age"`
	City      string `json:"city"`
	Accountid int64  `json:"accountid"`
}

/*
--------------------------------------------------------------------
----------------------- 2. Accounts --------------------------------
--------------------------------------------------------------------
*/

type Account struct {
	AccountID        int64  `json:"accountid"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	PhoneNumber      string `json:"phoneNumber"`
	DateOfBirth      string `json:"dateOfBirth"`
	SecurityQuestion string `json:"securityQuestion"`
	SecurityAnswer   string `json:"securityAnswer"`
}

type AccountSummary struct {
	AccountID   int64  `json:"accountid"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Email       string `json:"Email"`
	DateOfBirth string `json:"DateOfBirth"`
}

type AccountChange struct {
	AccountID        int64  `json:"accountid"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phoneNumber"`
	DateOfBirth      string `json:"dateOfBirth"`
	SecurityQuestion string `json:"securityQuestion"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
--------------------------------------------------------------------
----------------------- 3. Reviews ---------------------------------
--------------------------------------------------------------------
*/

type Review struct {
	ReviewID       int64   `json:"reviewID"`
	ProductName    string  `json:"productName"`
	Picture        string  `json:"productImage"`
	Score          int64   `json:"score"`
	BoughtFrom     string  `json:"locationBought"`
	BoughtFromLat  float64 `json:"boughtFromLat"`
	BoughtFromLong float64 `json:"boughtFromLong"`
	BoughtFor      float32 `json:"pricePaid"`
	Currency       string  `json:"currency"`
	FullReview     string  `json:"fullReview"`
	ProfileID      int64   `json:"createdBy"`
}

// type ParsedReview struct {
// 	ReviewID       int64  `json:"reviewID"`
// 	ProductName    string `json:"productName"`
// 	Picture        []byte `json:"productImage"`
// 	Score          string `json:"score"`
// 	BoughtFrom     string `json:"locationBought"`
// 	BoughtFromLat  string `json:"boughtFromLat"`
// 	BoughtFromLong string `json:"boughtFromLong"`
// 	BoughtFor      string `json:"pricePaid"`
// 	Currency       string `json:"currency"`
// 	FullReview     string `json:"fullReview"`
// 	ProfileID      string `json:"createdBy"`
// }

// ^^ Not currently needed ^^

type ReviewSummary struct {
	ReviewID    int64  `json:"reviewID"`
	ProductName string `json:"productName"`
	Picture     []byte `json:"productImage"`
	Score       int32  `json:"score"`
	ProfileID   int64  `json:"createdBy"`
}

type EditReview struct {
	FullReview string `json:"fullReview"`
}

type EditScore struct {
	Score int32 `json:"score"`
}
