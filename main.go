package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>testing</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "to get in touch please send email to <a href = \"mailto:wippyz@hotmail.com\">wippyz@hotmail.com </a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1> Page Not Found</h1><p>Please email us</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
