package com

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/reg.html")
}

func RegComplete(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	age := r.FormValue("age")
	fmt.Fprintf(w, "Имя %s, ему %s лет", name, age)
	// http.ServeFile(w, r, "static/reg_complete.html")
}

func Routers(port string) {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/auth", GetUser)
	http.HandleFunc("/reg_complete", RegComplete)
	http.ListenAndServe(port, nil)
}
