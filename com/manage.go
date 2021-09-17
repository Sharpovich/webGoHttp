package com

import (
	"net/http"
	"project/com/apps"
)

func Routers(port string) {
	http.HandleFunc("/", apps.HomePage)
	http.HandleFunc("/auth", apps.GetAuth)
	http.HandleFunc("/reg_complete", apps.IndexHandler)
	http.HandleFunc("/postform", apps.GetUser)

	http.ListenAndServe(port, nil)
}
