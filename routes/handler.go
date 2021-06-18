package route

import (
	"fmt"
	"test/database"
	"test/model"
	"test/service"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	database.InitDatabase()
	result := service.GetBooksService()
	return c.JSON(result)
}

func NewBook(c *fiber.Ctx) error {
	database.InitDatabase()
	book := new(model.Book)

	if err := c.BodyParser(book); err != nil {
		fmt.Println(err)
		return err
	}
	result := service.NewBookService(book)
	return c.JSON(result)
}

func UpdateBook(c *fiber.Ctx) error {
	database.InitDatabase()
	bookData := new(model.Book)
	c.BodyParser(bookData)
	id := c.Params("id")
	result := service.UpdateService(id, bookData)
	if result.Title == "" {
		return c.JSON("no book exist for the given id")
	} else {
		return c.JSON(result)
	}
}

func DeleteBook(c *fiber.Ctx) error {
	database.InitDatabase()
	id := c.Params("id")
	result := service.DeleteBookService(id)
	if result.Title == "" {
		return c.JSON("book does not exist with the given id")
	}
	return c.JSON(result)
}

func DataInsert(c *fiber.Ctx) error {
	database.InitDatabase()
	result := service.DataInsertService()
	if result {
		return c.Status(200).JSON("data inserted sucessfully")
	} else {
		return c.Status(200).JSON("failed to insert data")
	}
}

func GetBookById(c *fiber.Ctx) error {
	database.InitDatabase()
	id := c.Params("id")
	result := service.GetBookByIdService(id)
	if result.Title == "" {
		return c.JSON("no book exist for the given id")
	} else {
		return c.JSON(result)
	}
}

func Mailer(c *fiber.Ctx) error {

	result := service.MailerService()
	if result {
		return c.Status(200).JSON("mail sent sucessfully")
	} else {
		return c.Status(500).JSON("unable to sent mail")
	}
}

func Channel(c *fiber.Ctx) error {

	result := service.ChannelService()
	return c.Status(200).JSON(result)
}

func Registeration(c *fiber.Ctx) error {
	database.InitDatabase()
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(err)
	}
	result := service.RegisterationService(user)
	if result {
		return c.Status(200).JSON("user sucessfully created")
	} else {
		return c.Status(500).JSON("unable to create user")
	}
}

func Login(c *fiber.Ctx) error {
	database.InitDatabase()
	userData := new(model.User)
	c.BodyParser(userData)
	result := service.LoginService(userData)
	return c.Status(200).JSON(result)
}
