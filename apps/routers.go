package apps

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type application struct {
	logError *log.Logger
	logInfo  *log.Logger
}

func (app *application) ConnDB() *sql.DB {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_CONTAINER_NAME"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open(os.Getenv("DB_CONF"), connStr)
	if err != nil {
		panic(err)
	}
	app.logInfo.Printf("Connection opened to  %v\n",
		strings.ToUpper(os.Getenv("DB_NAME")))
	return db
}

func Routers(host, port string) {
	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	logError := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Инициализируем новую структуру с зависимостями приложения.
	app := &application{
		logError: logError,
		logInfo:  logInfo,
	}
	route := http.NewServeMux()
	// static files
	fs := http.FileServer(http.Dir("./static"))
	route.Handle("/static/", http.StripPrefix("/static/", fs))
	// home page
	route.HandleFunc("/", app.HomePage)
	// auth
	route.HandleFunc("/auth", app.GetAuth)
	route.HandleFunc("/auth/", app.GetAuth)
	route.HandleFunc("/auth/postform_authentication", app.GetUser)
	// review list users
	route.HandleFunc("/users", app.IndexHandler)

	logInfo.Printf("Connection web-server on %v%v\n", host, port)
	err := http.ListenAndServe(port, route)
	if err != nil {
		logError.Fatal(err)
	}
}
