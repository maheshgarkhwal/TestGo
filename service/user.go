package service

import (
	"fmt"
	"log"
	"os"
	"strings"
	"test/database"
	"test/model"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	fmt.Println("login module")

	userData := new(model.User)
	c.BodyParser(userData)
	db := database.DBConn
	var user model.User
	fmt.Println(userData)
	db.Find(&user, "Email = ?", userData.Email)
	fmt.Print(">>>>>", user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))

	if err != nil {
		fmt.Print("error in compareHash", err)
	} else {
		fmt.Print("else")
		token, err := CreateToken(user.Email)
		if err != nil {
			c.Status(500).JSON(err.Error())
			return
		}
		c.Status(200).JSON(token)

	}
}

func CreateToken(userId string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Authentication(c *fiber.Ctx) {
	tokenString := c.Get("Authorization")
	tokenString1 := strings.Split(tokenString, " ")

	hmacSecretString := os.Getenv("ACCESS_SECRET") // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenString1[1], func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		fmt.Println("error:", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		c.Status(200).JSON("authorize")

	} else {
		log.Println("Invalid JWT Token")
		c.Status(400).JSON("invalid json token")
	}
}