package apps

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	logInfo := log.New(os.Stdout, "INFO:\t",
		log.Ldate|log.Ltime)
	logError := log.New(os.Stderr, "ERROR:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

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
		logError.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		logError.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	logInfo.Printf("Page to /: %v\n", http.StatusOK)
}
