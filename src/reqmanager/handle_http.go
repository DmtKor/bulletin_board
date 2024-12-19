package reqmanager

import (
	"net/http"
)

func HandleHTTP(addr string) error {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/browse", browseHandler)
	http.HandleFunc("/examine", examineHandler)
	err := http.ListenAndServe(addr, nil)
	return err
}