package apps

import (
	"html/template"
	"net/http"
)

func (app *application) HomePage(w http.ResponseWriter, r *http.Request) {
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
		app.logError.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.logError.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	app.logInfo.Printf("Page to /: %v\n", http.StatusOK)
}
