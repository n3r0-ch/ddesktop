package main

import (
	"log"
	"ddesktop/server"
	"ddesktop/wsproxy"
	"net/http"
)

func main() {
	//Start server
	log.Printf("Starting server on http://0.0.0.0:9000...")
	r := server.NewRouter()
	http.Handle("/static/", server.HTTPProvider(http.StripPrefix("/static/", http.FileServer(http.Dir("webroot/static/"))), "GetStatic"))
	http.Handle("/websockify", wsproxy.WsProxy("localhost:6080"))
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":9000", nil))
}
