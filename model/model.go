package model

// Currently effectively an empty file, this is
// at the minute, bloat.

// however these structs will be used extensively to interact
// with a mySQL database

type Person struct {
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Age       int32  `json:"age"`
	City      string `json:"city"`
}

type Account struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	UserName         string `json:"userName"`
	PhoneNumber      string `json:"phoneNumber"`
	SecurityQuestion string `json:"securityQuestion"`
	SecurityAnswer   string `json:"securityAnswer"`
}

type Review struct {
	ProductName string  `json:"productName"`
	Picture     []byte  `json:"productImage"`
	Score       int32   `json:"score"`
	BoughtFrom  string  `json:"locationBought"`
	BoughtFor   float32 `json:"pricePaid"`
	FullReview  string  `json:"fullReview"`
}
