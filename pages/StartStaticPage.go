package pages

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path"
)

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
