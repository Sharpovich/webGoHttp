package apps

import (
	"database/sql"
	"html/template"
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

func (app *application) GetAuth(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"static/templates/authentication.html",
		"static/templates/base.html",
		"static/templates/footer.html",
		"static/templates/header.html",
	}
	if r.URL.Path != "/auth/" && r.URL.Path != "/auth" {
		http.Redirect(w, r, "/auth", http.StatusFound)
		app.logInfo.Printf("Redirecting to /auth: %v\n", http.StatusFound)
		return
	}
	if r.URL.Path == "/auth/" {
		http.Redirect(w, r, "/auth", http.StatusFound)
		app.logInfo.Printf("Redirecting to /auth: %v\n", http.StatusFound)
		return
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logError.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.logError.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	app.logInfo.Printf("Page to /auth: %v\n", http.StatusOK)
}

func (app *application) GetUser(w http.ResponseWriter, r *http.Request) {
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
		app.logInfo.Printf("Connection opened to  %v\n",
			strings.ToUpper(os.Getenv("DB_NAME")))
		defer db.Close()
		_, er := db.Exec("insert into users (firstname, lastname, city) values ($1, $2, $3)",
			firstname, lastname, city)
		if er != nil {
			panic(er)
		}
		app.logInfo.Printf("Data added to database %v\n",
			strings.ToUpper(os.Getenv("DB_NAME")))
		defer app.logInfo.Printf("Connection closed to  %v\n",
			strings.ToUpper(os.Getenv("DB_NAME")))
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
		app.logInfo.Printf("Page to /auth/postform_authentication: %v\n",
			http.StatusOK)
		return
	}
	http.Redirect(w, r, "/auth", http.StatusFound)
	app.logInfo.Printf("Redirecting to /auth: %v\n", http.StatusFound)
}
