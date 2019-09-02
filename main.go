package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter , r *http.Request){
	w.Header().Set("content-type","text/html")
	fmt.Fprint(w, "<h1>Welcome</h1>")

}
func contact(w http.ResponseWriter , r *http.Request){
	w.Header().Set("content-type","text/html")
	fmt.Fprint(w, "to get in touch please  send email to <a href = \"mailto:wippyz@hotmail.com\">wippyz@hotmail.com </a>.")

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/",home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
