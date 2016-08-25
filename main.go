package main

import (
	"html/template"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/resume", resume)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(res, nil)
}

func resume(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/resume.html")
	t.Execute(res, nil)
}
