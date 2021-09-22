package apps

import (
	"log"
	"net/http"
	"os"
)

type Application struct {
	logError *log.Logger
	logInfo  *log.Logger
}

func Logger() *Application {
	logInfo := log.New(os.Stdout, "INFO:\t",
		log.Ldate|log.Ltime)
	logError := log.New(os.Stderr, "ERROR:\t",
		log.Ldate|log.Ltime|log.Lshortfile)
	app := &Application{
		logError: logError,
		logInfo:  logInfo,
	}
	return app
}
func Routers(host, port string) {
	route := http.NewServeMux()
	// home page
	route.HandleFunc("/", HomePage)
	// auth
	route.HandleFunc("/auth", GetAuth)
	route.HandleFunc("/auth/", GetAuth)
	route.HandleFunc("/auth/postform_authentication", GetUser)
	// review list users
	route.HandleFunc("/users", IndexHandler)

	Logger().logInfo.Printf("Connection web-server on %v%v\n", host, port)
	err := http.ListenAndServe(port, route)
	if err != nil {
		Logger().logError.Fatal(err)
	}
}
