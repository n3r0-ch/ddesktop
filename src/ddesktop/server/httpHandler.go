package server

import (
	"net/http"
	"github.com/spf13/viper"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webroot/index.html")
}

func GetFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webroot/favicon.ico")
}

func RedirectHttps(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "https://" + r.Host + ":" + viper.GetString("server.port.https") + r.RequestURI, http.StatusMovedPermanently)
}