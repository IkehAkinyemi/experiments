package main

import (
	"html/template"
	"net/http"
	"time"
)

func formDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formDate}
	t:= template.New("tmpl.html").Funcs(funcMap)
	t, err := t.ParseFiles("tmpl.html", "t2.html")
	if err != nil {
		http.Error(w, "error occurred", http.StatusInternalServerError)
		return
	}
	t.Execute(w, time.Now())
}

func main() {
	s := http.Server{
		Addr: ":8900",
	}
	http.HandleFunc("/", process)
	s.ListenAndServe()
}