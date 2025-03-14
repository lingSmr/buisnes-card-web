package main

import (
	"log"
	"net/http"
	"path"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.Router{}
	routes(&r)
	log.Println("Server started on port :9090!")
	http.ListenAndServe(":9090", &r)
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	r.GET("/", StartPage)
	r.GET("/prikol", PrikolPage)
}

func StartPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := path.Join("public", "html", "StartPage.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PrikolPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := path.Join("public", "html", "prikol.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
