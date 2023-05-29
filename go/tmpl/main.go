package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t := template.New("tmpl.html")
	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, template.HTML(r.FormValue("comment")))
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