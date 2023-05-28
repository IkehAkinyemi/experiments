package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	// t, _ := template.ParseGlob("*.html")
	rand.NewSource(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	s := http.Server{
		Addr: ":8900",
	}
	http.HandleFunc("/", process)
	s.ListenAndServe()
}