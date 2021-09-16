package com

import (
	"net/http"
	"project/com/apps"
)

func Routers(port string) {
	http.HandleFunc("/", apps.HomePage)
	http.HandleFunc("/auth", apps.GetUser)
	http.HandleFunc("/reg_complete", apps.RegComplete)

	http.ListenAndServe(port, nil)
}
