package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	rand.NewSource(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("tmpl.html", "red_hello.html")
	} else {
		t, _ = template.ParseFiles("tmpl.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	s := http.Server{
		Addr: ":8900",
	}
	http.HandleFunc("/process", process)
	s.ListenAndServe()
}