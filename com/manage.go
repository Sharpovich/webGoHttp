package com

import (
	"fmt"
	"log"
	"net/http"
	"project/com/apps"
)

func Routers(host, port string) {
	route := http.NewServeMux()
	// StaticFile
	fileServer := http.FileServer(http.Dir("./static/"))
	route.Handle("/static/", http.StripPrefix("/static", fileServer))
	// home page
	route.HandleFunc("/", apps.HomePage)
	// auth
	route.HandleFunc("/auth", apps.GetAuth)
	route.HandleFunc("/auth/", apps.GetAuth)
	route.HandleFunc("/auth/postform_authentication", apps.GetUser)
	// review list users
	route.HandleFunc("/users", apps.IndexHandler)

	fmt.Printf("Connection web-server on %v%v\n", host, port)
	err := http.ListenAndServe(port, route)
	if err != nil {
		log.Fatal(err)
	}
}
