package apps

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type Users struct {
	Id        int
	Firstname string
	Lastname  string
	City      string
}

func (app *application) IndexHandler(w http.ResponseWriter, r *http.Request) {

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		"db",
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open(os.Getenv("DB_CONF"), connStr)
	if err != nil {
		panic(err)
	}
	app.logInfo.Printf("Connection opened to  %v\n",
		strings.ToUpper(os.Getenv("DB_NAME")))
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
	app.logInfo.Printf("Data from database %v loaded\n",
		strings.ToUpper(os.Getenv("DB_NAME")))
	app.logInfo.Printf("Connection closed to  %v\n",
		strings.ToUpper(os.Getenv("DB_NAME")))
	files := []string{
		"static/templates/list_users.html",
		"static/templates/base.html",
		"static/templates/footer.html",
		"static/templates/header.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.logError.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	err = tmpl.Execute(w, users)
	if err != nil {
		app.logError.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	defer app.logInfo.Printf("Page to /users: %v\n", http.StatusOK)
}
