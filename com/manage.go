package com

import (
	"fmt"
	"log"
	"net/http"
	"project/com/apps"
)

func Routers(host, port string) {
	// home page
	http.HandleFunc("/", apps.HomePage)
	// auth
	http.HandleFunc("/auth", apps.GetAuth)
	http.HandleFunc("/auth/", apps.GetAuth)
	http.HandleFunc("/auth/postform_authentication", apps.GetUser)
	// review list users
	http.HandleFunc("/users", apps.IndexHandler)

	fmt.Printf("Connection web-server on %v%v\n", host, port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
