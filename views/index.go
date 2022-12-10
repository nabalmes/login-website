package views

import (
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	activeSession := GetActiveSession(r)
	if activeSession == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data["Title"] = "Index"
	tmpl.Execute(w, data)
}
