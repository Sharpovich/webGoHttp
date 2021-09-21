package apps

import (
	"html/template"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"static/templates/home.page.tmpl",
		"static/templates/base.tmpl",
		"static/templates/footer.tmpl",
		"static/templates/header.tmpl",
	}
	filesNotFound := []string{
		"static/templates/not_found.tmpl",
		"static/templates/base.tmpl",
		"static/templates/footer.tmpl",
		"static/templates/header.tmpl",
	}
	if r.URL.Path != "/" {
		nf, _ := template.ParseFiles(filesNotFound...)
		nf.Execute(w, nil)
		return
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
