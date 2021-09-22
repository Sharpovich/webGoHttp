package apps

import (
	"html/template"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"static/templates/home.page.html",
		"static/templates/base.html",
		"static/templates/footer.html",
		"static/templates/header.html",
	}
	filesNotFound := []string{
		"static/templates/not_found.html",
		"static/templates/base.html",
		"static/templates/footer.html",
		"static/templates/header.html",
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
