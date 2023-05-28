package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, r.FormValue("comment"))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}

func main() {
	s := http.Server{
		Addr: ":8900",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)
	s.ListenAndServe()
}