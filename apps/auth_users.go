package apps

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type User struct {
	Firstname string
	Lastname  string
	City      string
}

func GetAuth(w http.ResponseWriter, r *http.Request) {
	logInfo := log.New(os.Stdout, "INFO:\t",
		log.Ldate|log.Ltime)
	logError := log.New(os.Stderr, "ERROR:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	files := []string{
		"static/templates/authentication.html",
		"static/templates/base.html",
		"static/templates/footer.html",
		"static/templates/header.html",
	}
	if r.URL.Path != "/auth/" && r.URL.Path != "/auth" {
		http.Redirect(w, r, "/auth", http.StatusFound)
		logInfo.Printf("Redirecting to /auth: %v\n", http.StatusFound)
		return
	}
	if r.URL.Path == "/auth/" {
		http.Redirect(w, r, "/auth", http.StatusFound)
		logInfo.Printf("Redirecting to /auth: %v\n", http.StatusFound)
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
	logInfo.Printf("Page to /auth: %v\n", http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	logInfo := log.New(os.Stdout, "INFO:\t",
		log.Ldate|log.Ltime)
	logError := log.New(os.Stderr, "ERROR:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	city := r.FormValue("city")

	files := []string{
		"static/templates/postform_authentication.html",
		"static/templates/base.html",
		"static/templates/footer.html",
		"static/templates/header.html",
	}
	users := []User{}
	p := User{Firstname: firstname, Lastname: lastname, City: city}
	users = append(users, p)

	if firstname != "" && lastname != "" && city != "" {
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
		_, er := db.Exec("insert into users (firstname, lastname, city) values ($1, $2, $3)",
			firstname, lastname, city)
		if er != nil {
			panic(er)
		}
		logInfo.Printf("Data added to database %v\n",
			strings.ToUpper(os.Getenv("DB_NAME")))
		defer logInfo.Printf("Connection closed to  %v\n",
			strings.ToUpper(os.Getenv("DB_NAME")))
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
		logInfo.Printf("Page to /auth/postform_authentication: %v\n",
			http.StatusOK)
		return
	}
	http.Redirect(w, r, "/auth", http.StatusFound)
	logInfo.Printf("Redirecting to /auth: %v\n", http.StatusFound)
}
