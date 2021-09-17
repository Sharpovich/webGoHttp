package apps

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Users struct {
	Id        int
	Firstname string
	Lastname  string
	City      string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=admin password=admin dbname=project sslmode=disable"
	db, err := sql.Open("postgres", connStr)
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

	tmpl, _ := template.ParseFiles("static/templates/reg_complete.html")
	tmpl.Execute(w, users)
}
