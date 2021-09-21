package com

import (
	"fmt"
	"log"
	"net/http"
	"project/com/apps"
)

// func Routers(host, port string) {
// 	// StaticFile
// 	fileServer := http.FileServer(http.Dir("./static/"))
// 	http.Handle("/static/", http.StripPrefix("/static", fileServer))
// 	// home page
// 	http.HandleFunc("/", apps.HomePage)
// 	// auth
// 	http.HandleFunc("/auth", apps.GetAuth)
// 	http.HandleFunc("/auth/", apps.GetAuth)
// 	http.HandleFunc("/auth/postform_authentication", apps.GetUser)
// 	// review list users
// 	http.HandleFunc("/users", apps.IndexHandler)

// 	fmt.Printf("Connection web-server on %v%v\n", host, port)
// 	err := http.ListenAndServe(port, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func Routers(host, port string) {
	mux := http.NewServeMux()
	// StaticFile
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// home page
	mux.HandleFunc("/", apps.HomePage)
	// auth
	mux.HandleFunc("/auth", apps.GetAuth)
	mux.HandleFunc("/auth/", apps.GetAuth)
	mux.HandleFunc("/auth/postform_authentication", apps.GetUser)
	// review list users
	mux.HandleFunc("/users", apps.IndexHandler)

	fmt.Printf("Connection web-server on %v%v\n", host, port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
