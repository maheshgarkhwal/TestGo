package service

import (
	"crypto/tls"
	"fmt"

	"github.com/gofiber/fiber"
	gomail "gopkg.in/mail.v2"
)

func Mailer(c *fiber.Ctx) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "mahesh.garkhwal@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "mukesh.garkhwal@gmail.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "mahesh.garkhwal@gmail.com", "unknown")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		c.Status(200).JSON("mail sent sucessfully")
	}
}
