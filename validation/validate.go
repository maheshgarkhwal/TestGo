package validate

import (
	"net/url"
	"regexp"
	"test/model"
)

var regexpEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ValidBook(a model.Book) url.Values {

	errs := url.Values{}
	// check if the title empty
	if a.Title == "" {
		errs.Add("title", "The title is required!")
	}
	// check the name field is between 3 to 120 chars
	if a.Author == "" {
		errs.Add("author", "The author is required!")
	}
	if a.Rating >= 6 || a.Rating <= 0 {
		errs.Add("rating", "The rating should in between 1 to 5!")
	}

	return errs
}

func ValidUser(u model.User) url.Values {

	errs := url.Values{}

	if u.Username == "" {
		errs.Add("username", "The username is required!")
	}

	if len(u.FirstName) <= 0 || len(u.FirstName) > 10 {
		errs.Add("firstname", "The firstname field must be between 0-10 chars!")
	}
	if u.Email == "" {
		errs.Add("email", "The email field is required!")
	}

	if !regexpEmail.MatchString(u.Email) {
		errs.Add("email", "The email field should be a valid email address!")
	}

	if len(u.Password) < 5 {
		errs.Add("Password", "The Password length should be greater than 5!")
	}

	return errs
}
