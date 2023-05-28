package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t:= template.New("tmpl.html")
	t, err := t.ParseFiles("tmpl.html", "t2.html")
	if err != nil {
		http.Error(w, "error occurred", http.StatusInternalServerError)
		return
	}
	c := `I asked: <i>"What's up?"</i>`
	t.Execute(w, c)
}

func main() {
	s := http.Server{
		Addr: ":8900",
	}
	http.HandleFunc("/", process)
	s.ListenAndServe()
}