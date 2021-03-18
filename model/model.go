package model

// Currently effectively an empty file, this is
// at the minute, bloat.

// however these structs will be used extensively to interact
// with a mySQL database

/*
	There will need to be 1/2/3 or more GORMs to interact with databases
	As not all info will be relevant for each request

	Work as so: Stringify => Concat => Encode response body

	Example
		FULL REVIEW GORM
		Profile fields taken from foreign key info
		All review info, read all in structs below.

		SHORT REVIEW GORM
		UserName
		Object
		Picture
		Score

	These will simply be used to map data easily into the desired
	formats with directly relevant information.

	Few ideas, we'll see...
*/

type Profile struct {
	UserName string `json:"userName"`
	Age      int32  `json:"age"`
	City     string `json:"city"`
}

type Account struct {
	FirstName        string `json:"firstName"`
	Lastname         string `json:"lastName"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	PhoneNumber      string `json:"phoneNumber"`
	DateOfBirth      string `json:"dateOfBirth"`
	SecurityQuestion string `json:"securityQuestion"`
	SecurityAnswer   string `json:"securityAnswer"`
	ProfileID        string `json:"linkedProfile"`
}

type Review struct {
	ReviewID    int64   `json:"reviewID"`
	ProductName string  `json:"productName"`
	Picture     []byte  `json:"productImage"`
	Score       int32   `json:"score"`
	BoughtFrom  string  `json:"locationBought"`
	BoughtFor   float32 `json:"pricePaid"`
	FullReview  string  `json:"fullReview"`
	ProfileID   int64   `json:"createdBy"`
}

type ReviewSummary struct {
	ReviewID    int64  `json:"reviewID"`
	ProductName string `json:"productName"`
	Picture     []byte `json:"productImage"`
	Score       int32  `json:"score"`
	ProfileID   int64  `json:"createdBy"`
}
