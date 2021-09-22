package apps

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	logInfo := log.New(os.Stdout, "INFO:\t",
		log.Ldate|log.Ltime)
	logError := log.New(os.Stderr, "ERROR:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	connStr := "user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	db, err := sql.Open(os.Getenv("DB_CONF"), connStr)
	if err != nil {
		panic(err)
	}
	logInfo.Printf("Connection opened to  %v\n",
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
	logInfo.Printf("Data from database %v loaded\n",
		strings.ToUpper(os.Getenv("DB_NAME")))
	logInfo.Printf("Connection closed to  %v\n",
		strings.ToUpper(os.Getenv("DB_NAME")))
	files := []string{
		"static/templates/list_users.html",
		"static/templates/base.html",
		"static/templates/footer.html",
		"static/templates/header.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		logError.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	err = tmpl.Execute(w, users)
	if err != nil {
		logError.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	defer logInfo.Printf("Page to /users: %v\n", http.StatusOK)
}
