package com

import (
	"log"
	"net/http"
	"project/com/apps"
)

func Routers(port string) {
	// home page
	http.HandleFunc("/", apps.HomePage)
	// auth
	http.HandleFunc("/auth", apps.GetAuth)
	http.HandleFunc("/auth/postform_authentication", apps.GetUser)
	// review list users
	http.HandleFunc("/users", apps.IndexHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
