package server

import (
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webroot/index.html")
}

func GetFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webroot/favicon.ico")
}

