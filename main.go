package main

import (
	"fmt"
	"test/database"
	"test/model"
	route "test/routes"

	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	dsn := "mahesh:Mahesh@g7@tcp(localhost:3306)/crud?parseTime=true"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connection Opened to Database")
	}
	database.DBConn.AutoMigrate(&model.Book{}, &model.Info{})
}

func main() {
	app := fiber.New()
	initDatabase()
	route.SetupRoutes(app)
	app.Listen(3000)
}
