package service

import (
	"fmt"
	"strconv"
	"test/database"
	"test/model"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gofiber/fiber"
)

func DataInsert(c *fiber.Ctx) {
	go D1()
	go D2()

	time.Sleep(10 * time.Second)
	c.Status(200).JSON("Data insert sucessfully")
}

func D1() {

	detail := new(model.Details)
	db := database.DBConn

	f, err := excelize.OpenFile("sample.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := f.GetRows("Sheet1")

	if err != nil {
		fmt.Print(err.Error())
	}

	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>", len(rows))

	for i := 1; i < len(rows)/2; i++ {

		for j := 0; j < 7; j++ {
			fmt.Print(rows[i][j], "\t")
			if j == 0 {
				if val, err := strconv.Atoi(rows[i][j]); err == nil {
					detail.Id = val
				}
			}

			if j == 1 {
				detail.FirstName = rows[i][j]
			}

			if j == 2 {
				detail.LastName = rows[i][j]
			}

			if j == 3 {
				detail.Gender = rows[i][j]
			}

			if j == 4 {
				detail.Country = rows[i][j]
			}

			if j == 5 {
				if val1, err := strconv.Atoi(rows[i][j]); err == nil {
					detail.Age = val1
				}
			}
			if j == 6 {
				detail.Date = rows[i][j]
			}
		}
		db.Create(&detail)
		fmt.Println()
	}
}

func D2() {

	detail := new(model.Details)
	db := database.DBConn

	f, err := excelize.OpenFile("sample.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := f.GetRows("Sheet1")

	if err != nil {
		fmt.Print(err.Error())
	}

	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>", len(rows))

	for i := len(rows) / 2; i < len(rows); i++ {

		for j := 0; j < 7; j++ {
			fmt.Print(rows[i][j], "\t")
			if j == 0 {
				if val, err := strconv.Atoi(rows[i][j]); err == nil {
					detail.Id = val
				}
			}

			if j == 1 {
				detail.FirstName = rows[i][j]
			}

			if j == 2 {
				detail.LastName = rows[i][j]
			}

			if j == 3 {
				detail.Gender = rows[i][j]
			}

			if j == 4 {
				detail.Country = rows[i][j]
			}

			if j == 5 {
				if val1, err := strconv.Atoi(rows[i][j]); err == nil {
					detail.Age = val1
				}
			}
			if j == 6 {
				detail.Date = rows[i][j]
			}
		}
		db.Create(&detail)
		fmt.Println()
	}
}
