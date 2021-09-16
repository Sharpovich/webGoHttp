package apps

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/templates/reg.html")
	firstname := r.URL.Query().Get("firstname")
	lastname := r.URL.Query().Get("lastname")
	city := r.URL.Query().Get("city")

	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(city)
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
