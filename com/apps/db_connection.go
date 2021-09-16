package apps

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Product struct {
	Id      int
	Model   string
	Company string
	Price   int
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=admin password=admin dbname=project sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []Product{}

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	tmpl, _ := template.ParseFiles("static/templates/array_users.html")
	tmpl.Execute(w, products)
}
