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