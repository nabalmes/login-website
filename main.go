package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/nabalmes/practice/api"
	"github.com/nabalmes/practice/models"
	"github.com/nabalmes/practice/views"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	BindIP = "0.0.0.0"
	Port   = ":4200"
)

func main() {
	fmt.Printf("Go to port System: %v%v/\n", BindIP, Port)
	CreateDB("practice")
	MigrateDB()

	RegisterHandlers()

	http.ListenAndServe(Port, nil)
}

func RegisterHandlers() {
	http.Handle("/templates", http.StripPrefix("/templates", http.FileServer(http.Dir("./templates"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", views.LoginHandler)
	http.HandleFunc("/api/", api.APIHandler)

}

func CreateDB(name string) *sql.DB {
	fmt.Println("Database Created")
	db, err := sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func MigrateDB() {
	fmt.Println("Database Migrated")
	user := models.User{}

	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/practice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&user)

}
