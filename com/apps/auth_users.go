package apps

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func GetAuth(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/templates/authentication.html")
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
