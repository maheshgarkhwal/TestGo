package database

import (
	"fmt"
	"test/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() {
	var err error
	dsn := "mahesh:Mahesh@g7@tcp(localhost:3306)/crud?parseTime=true"
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connection Opened to Database")
	}
	DBConn.AutoMigrate(&model.Student{}, &model.Book{}, &model.Info{}, &model.User{})
}
