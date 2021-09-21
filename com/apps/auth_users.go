package apps

import (
	"database/sql"
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

type User struct {
	Firstname string
	Lastname  string
	City      string
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	city := r.FormValue("city")

	files := []string{
		"static/templates/postform_authentication.tmpl",
		"static/templates/base.tmpl",
		"static/templates/footer.tmpl",
		"static/templates/header.tmpl",
	}
	users := []User{}
	p := User{Firstname: firstname, Lastname: lastname, City: city}
	users = append(users, p)
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	err = tmpl.Execute(w, users)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

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
