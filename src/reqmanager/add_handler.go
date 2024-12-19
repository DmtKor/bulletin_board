package reqmanager

import "net/http"

func addHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/add.html")
}