package apps

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func GetAuth(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"static/templates/authentication.tmpl",
		"static/templates/base.tmpl",
		"static/templates/footer.tmpl",
		"static/templates/header.tmpl",
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	city := r.FormValue("city")
	fmt.Fprintf(w, "Имя: %s Фамилия: %s Город: %s", firstname, lastname, city)

	connStr := "user=admin password=admin dbname=project sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, er := db.Exec("insert into users (firstname, lastname, city) values ($1, $2, $3)",
		firstname, lastname, city)
	if er != nil {
		panic(er)
	}
}
