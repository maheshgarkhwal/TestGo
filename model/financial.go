package model

type Details struct {
	Id        int    `json:"id"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Gender    string `json:"gender"`
	Country   string `json:"country"`
	Age       int    `json:"age"`
	Date      string `json:"date"`
}
