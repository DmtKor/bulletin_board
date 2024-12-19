package reqmanager

import (
	"bulletin_board/src/dbmanager"
	"html/template"
	"net/http"
	"strconv"
)


func browseHandler(w http.ResponseWriter, r *http.Request) {
    page, err := strconv.Atoi(r.URL.Query().Get("page"))
	// Wrong page
	if err != nil {
		page = 1
	}
	data := dbmanager.GetBrowseElementsByPage(uint(page))
	if len(data.Data) == 0 {
		tmpl, _ := template.ParseFiles("static/templates/error.html")
		tmpl.Execute(w, "Не найдено результатов по запросу")
		return
	}

	tmpl, _ := template.ParseFiles("static/templates/browse.html")
	tmpl.Execute(w, data)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/add.html")
}

type examineView struct {
	Title string
	Num   string
	Text  string
}

func examineHandler(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(r.URL.Query().Get("num"))
	// Wrong page
	if err != nil {
		tmpl, _ := template.ParseFiles("static/templates/error.html")
		tmpl.Execute(w, err)
		return
	}
	data, err := dbmanager.GetElement(uint(num))
	if err != nil {
		tmpl, _ := template.ParseFiles("static/templates/error.html")
		tmpl.Execute(w, err)
		return
	}
	view := examineView { Num: strconv.Itoa(num + 1), Title: data.Title, Text: data.Text }
	tmpl, _ := template.ParseFiles("static/templates/examine.html")
	tmpl.Execute(w, view)
}

func HandleHTTP(addr string) error {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/browse", browseHandler)
	http.HandleFunc("/examine", examineHandler)
	err := http.ListenAndServe(addr, nil)
	return err
}