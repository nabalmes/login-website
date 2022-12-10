package api

import (
	"fmt"
	"net/http"

	"github.com/nabalmes/practice/models"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	// activeSession := GetActiveSession(r)

	// if activeSession != "" {
		fmt.Println("test")
		fname := r.FormValue("firstname")
		lname := r.FormValue("lastname")
		username := r.FormValue("username")
		password := r.FormValue("password")

		user := models.User{}

		user.FirstName = fname
		user.LastName = lname
		user.Username = username
		user.Password = hashPassword(password)
		user.Save()

		res := map[string]interface{}{
			"status": "ok",
		}
		ReturnJSON(w, r, res)
	}
	// } else {
	// 	res := map[string]interface{}{
	// 		"status":     "error",
	// 		"permission": "denied",
	// 	}
	// 	ReturnJSON(w, r, res)
	// }
// }

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

// func GetActiveSession(r *http.Request) string {
// 	key, err := r.Cookie("session")
// 	if err == nil && key != nil {
// 		return key.Value
// 	}
// 	return ""
// }
