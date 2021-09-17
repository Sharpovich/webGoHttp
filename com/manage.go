package com

import (
	"net/http"
	"project/com/apps"
)

func Routers(port string) {
	http.HandleFunc("/", apps.HomePage)
	http.HandleFunc("/auth", apps.GetAuth)
	http.HandleFunc("/users", apps.IndexHandler)
	http.HandleFunc("/postform_authentication", apps.GetUser)

	http.ListenAndServe(port, nil)
}
