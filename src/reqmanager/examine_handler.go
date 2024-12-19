package reqmanager

import (
	"bulletin_board/src/dbmanager"
	"html/template"
	"net/http"
	"strconv"
)

type examineView struct {
	Title    string
	Num      string
	Text     string
	BackPage string
}

func examineHandler(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(r.URL.Query().Get("num"))
	// Wrong num
	if err != nil {
		tmpl, _ := template.ParseFiles("static/templates/error.html")
		tmpl.Execute(w, err)
		return
	}
	backpage := r.URL.Query().Get("backpage")
	// Wrong page
	if backpage == "" {
		backpage = "1"
	}
	data, err := dbmanager.GetElement(uint(num))
	if err != nil {
		tmpl, _ := template.ParseFiles("static/templates/error.html")
		tmpl.Execute(w, err)
		return
	}
	view := examineView { Num: strconv.Itoa(num + 1), Title: data.Title, Text: data.Text, BackPage: backpage }
	tmpl, _ := template.ParseFiles("static/templates/examine.html")
	tmpl.Execute(w, view)
}