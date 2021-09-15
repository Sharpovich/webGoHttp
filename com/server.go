package com

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/templates/index.html")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/templates/reg.html")
}

type Data struct {
	Name string
	Age  string
}

func RegComplete(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Name: r.FormValue("username"),
		Age:  r.FormValue("age"),
	}
	tmpl, _ := template.ParseFiles("static/templates/reg_complete.html")
	tmpl.Execute(w, data)

}

func Routers(port string) {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/auth", GetUser)
	http.HandleFunc("/reg_complete", RegComplete)

	http.ListenAndServe(port, nil)
}
