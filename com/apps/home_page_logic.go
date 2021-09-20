package apps

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.ServeFile(w, r, "static/templates/not_found.html")
		return
	}
	http.ServeFile(w, r, "static/templates/index.html")
}
