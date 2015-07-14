package server

import (
	"net/http"
	"github.com/spf13/viper"
	"fmt"
	"log"
)

func RedirectHttps(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "https://" + r.Host + ":" + viper.GetString("server.port.https") + r.RequestURI, http.StatusMovedPermanently)
}

func Static() http.HandlerFunc {
    fs := http.FileServer(http.Dir("webroot"))
    realHandler := http.StripPrefix("", fs).ServeHTTP
    return func(w http.ResponseWriter, req *http.Request) {

		//Cache the next 30 days
		w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d, public, must-revalidate, proxy-revalidate", 2592000))

        realHandler(w, req)

        //Log
		log.Printf(
			"%s\t%s\t",
			req.Method,
			req.RequestURI,
		)
    }
}
