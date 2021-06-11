package service

import (
	"fmt"
	"test/database"
	"test/model"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func Registeration(c *fiber.Ctx) {
	db := database.DBConn

	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
		return
	}
	password := []byte(user.Password)

	fmt.Print(password)
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))
	user.Password = string(hashedPassword)
	fmt.Print(">>>user info  \n", user)

	db.Create(&user)
	c.Status(200).JSON("user sucessfully created")

}

func Login(c *fiber.Ctx) {
	fmt.Print("login module")
}
