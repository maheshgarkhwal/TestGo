package model

type Info struct {
	Id        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Gender    string
	Country   string
	Age       int
	Date      string
}
