package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	d := []string{"Mon", "Tue", "Wed", "Thurs", "Fri", "Sat", "Sun"}
	t.Execute(w, d)
}

func main() {
	s := http.Server{
		Addr: ":8900",
	}
	http.HandleFunc("/", process)
	s.ListenAndServe()
}