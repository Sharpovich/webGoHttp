package apps

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Users struct {
	Id        int
	Firstname string
	Lastname  string
	City      string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"
	db, err := sql.Open(os.Getenv("DB_CONF"), connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := []Users{}

	for rows.Next() {
		p := Users{}
		err := rows.Scan(&p.Id, &p.Firstname, &p.Lastname, &p.City)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	files := []string{
		"static/templates/list_users.tmpl",
		"static/templates/base.tmpl",
		"static/templates/footer.tmpl",
		"static/templates/header.tmpl",
	}
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
}
