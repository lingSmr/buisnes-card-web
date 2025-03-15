package main

import (
	"log"
	"net/http"
	"simple_visit_site/pages"

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
	r.GET("/", pages.StartPage)
	r.GET("/prikol", pages.PrikolPage)
	r.GET("/photo/sanya", pages.SanyaPage)
	r.GET("/photo/david", pages.DavidPage)
	r.GET("/photo/green", pages.GreenPage)
}
