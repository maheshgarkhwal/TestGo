package route

import (
	"fmt"
	inter "test/interface"
	"test/model"
	"test/service"

	"github.com/gofiber/fiber/v2"
)

//Getting all Books
func GetBooks(c *fiber.Ctx) error {
	result := Get(Service)
	return c.Status(200).JSON(fiber.Map{"message": "All books", "result": result})
}

func Get(s inter.IBook) []model.Book {
	result := s.GetBookService()
	return result
}

//registering a new book
func NewBook(c *fiber.Ctx) error {
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		fmt.Println(err)
		return err
	}
	result := BookRegister(*book, Service)
	return c.Status(200).JSON(fiber.Map{"message": "Book added sucessfully", "result": result})
}

func BookRegister(b model.Book, s inter.IBook) model.Book {
	result := s.NewBookService(b)
	return result
}

//updating a book
func UpdateBook(c *fiber.Ctx) error {
	bookData := new(model.Book)
	c.BodyParser(bookData)
	id := c.Params("id")
	result := BookUpdate(Service, *bookData, id)
	if result.Title == "" {
		return c.Status(400).JSON("no book exist for the given id")
	} else {
		return c.Status(200).JSON(fiber.Map{"message": "Book updated sucessfully", "result": result})
	}
}

func BookUpdate(s inter.IBook, bookData model.Book, id string) model.Book {
	result := s.UpdateBookService(bookData, id)
	return result
}

//delete book
func DeleteBook(c *fiber.Ctx) error {

	id := c.Params("id")
	result := BookDelete(Service, id)
	if result.Title == "" {
		return c.Status(400).JSON("book does not exist with the given id")
	}
	return c.Status(200).JSON(fiber.Map{"message": "Book deleted sucessfully", "result": result})
}

func BookDelete(s inter.IBook, id string) model.Book {
	result := s.DeleteBookService(id)
	return result
}

func GetBookById(c *fiber.Ctx) error {

	id := c.Params("id")
	result := GetBook(Service, id)
	if result.Title == "" {
		return c.JSON(fiber.Map{"status": "true", "message": "no book exist for the given id"})
	} else {
		return c.JSON(fiber.Map{"status": "true", "result": result})
	}
}

func GetBook(s inter.IBook, id string) model.Book {
	result := s.GetBookByIdService(id)
	return result
}

//data insertion service
func DataInsert(c *fiber.Ctx) error {

	result := service.DataInsertService()
	if result {
		return c.Status(200).JSON("data inserted sucessfully")
	} else {
		return c.Status(200).JSON("failed to insert data")
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

	userData := new(model.User)
	c.BodyParser(userData)
	result := service.LoginService(userData)
	return c.Status(200).JSON(result)
}

func GetUser(c *fiber.Ctx) error {
	pg := c.Query("page")
	limit := c.Query("limit")
	result, message := service.GetUserService(pg, limit)
	if message != "" {
		return c.JSON(fiber.Map{"status": "failed", "result": message})
	} else {
		if len(result) < 1 {
			return c.JSON(fiber.Map{"status": "sucess", "result": "no more record left"})
		} else {
			return c.JSON(fiber.Map{"status": "sucess", "result": result})
		}
	}
}

//Interface

func PostStudent(c *fiber.Ctx) error {
	a := inter.Student{
		Name:   "arjunartistic",
		School: "asian school",
		Class:  "4th",
		RollNo: "21",
	}
	register(a)

	return c.JSON(">")
}

func register(s inter.StudentRegister) inter.Student {
	result := s.CreateStudent()
	return result
}
