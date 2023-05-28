package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	// t, _ := template.ParseGlob("*.html")
	t.Execute(w, "Hello word")
}

func main() {
	s := http.Server{
		Addr: ":8900",
	}
	http.HandleFunc("/", process)
	s.ListenAndServe()
}