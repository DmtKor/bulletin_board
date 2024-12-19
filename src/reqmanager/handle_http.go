package reqmanager

import (
	"net/http"
)

func HandleHTTP(addr string) error {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/browse", browseHandler)
	http.HandleFunc("/examine", examineHandler)
	http.HandleFunc("/new", newHandler)
	err := http.ListenAndServe(addr, nil)
	return err
}