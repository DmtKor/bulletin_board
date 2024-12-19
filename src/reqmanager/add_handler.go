package reqmanager

import (
	"bulletin_board/src/dbmanager"
	"html/template"
	"net/http"
	"strconv"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/add.html")
}

// address "/new..." is called from form on "/add" page
func newHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
    text := r.FormValue("text")
	i, err := dbmanager.AddElement(dbmanager.Element{ Title: title, Text: text })
	if err != nil {
		tmpl, _ := template.ParseFiles("static/templates/error.html")
		tmpl.Execute(w, err)
		return
	}
	http.Redirect(w, r, "/examine?num=" + strconv.Itoa(i), http.StatusSeeOther)
}