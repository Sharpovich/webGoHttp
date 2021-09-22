package apps

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	logError *log.Logger
	logInfo  *log.Logger
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
