package inter

import (
	"fmt"
	"test/database"
)

type Student struct {
	Name   string
	School string
	Class  string
	RollNo string
}

func (st Student) CreateStudent() Student {
	db := database.DBConn
	db.Create(&st)
	fmt.Print(st)
	return st
}

type StudentRegister interface {
	CreateStudent() Student
}
