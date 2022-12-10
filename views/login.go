package views

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
	"github.com/nabalmes/practice/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	tmpl := template.Must(template.ParseFiles("./templates/login.html"))

	activeSession := GetActiveSession(r)

	if activeSession != "" {
		http.Redirect(w, r, "/balmes/landingpage", http.StatusSeeOther)
	}

	if r.Method == "POST" {
		dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/practice?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println("Failed to Connect to the Database", err)
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		user := models.User{}

		db.Where("username = ?", username).Find(&user)

		if CheckPasswordHash(password, user.Password) {
			newSession := uuid.NewString()

			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "Session",
				Value: newSession,
			})

			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "username",
				Value: user.Username,
			})

			http.Redirect(w, r, "/practice/test", http.StatusSeeOther)
		}
	}

	data["Title"] = "Login"
	tmpl.Execute(w, data)
}

func GetActiveSession(r *http.Request) string {
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key.Value
	}
	return ""
}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
