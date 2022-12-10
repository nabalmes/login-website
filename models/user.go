package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Username  string
	Password  string
}

func (e *User) Save() {
	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/practice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to Connect to the Database ", err)
	}

	db.Save(e)
}
