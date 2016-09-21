package main

import (
	"html/template"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/en", index_en)
	http.HandleFunc("/resume", resume)
	http.HandleFunc("/resume_en", resume_en)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(res, nil)
}

func index_en(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index_en.html")
	t.Execute(res, nil)
}

func resume(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/resume.html")
	t.Execute(res, nil)
}

func resume_en(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/resume_en.html")
	t.Execute(res, nil)
}
