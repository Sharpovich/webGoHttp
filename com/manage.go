package com

import (
	"log"
	"net/http"
	"os"
	"project/com/apps"
)

func Routers(host, port string) {
	logInfo := log.New(os.Stdout, "INFO:\t",
		log.Ldate|log.Ltime)
	logError := log.New(os.Stderr, "ERROR:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	route := http.NewServeMux()
	// home page
	route.HandleFunc("/", apps.HomePage)
	// auth
	route.HandleFunc("/auth", apps.GetAuth)
	route.HandleFunc("/auth/", apps.GetAuth)
	route.HandleFunc("/auth/postform_authentication", apps.GetUser)
	// review list users
	route.HandleFunc("/users", apps.IndexHandler)

	logInfo.Printf("Connection web-server on %v%v\n", host, port)
	err := http.ListenAndServe(port, route)
	if err != nil {
		logError.Fatal(err)
	}
}
