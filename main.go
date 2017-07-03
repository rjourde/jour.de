package hello

import (
    "html/template"
	"net/http"
)

func init() {

    http.Handle("/static/", http.FileServer(http.Dir(".")))
    
    http.HandleFunc("/", index)
	http.HandleFunc("/en", index_en)
	http.HandleFunc("/resume", resume)
	http.HandleFunc("/resume_en", resume_en)

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
