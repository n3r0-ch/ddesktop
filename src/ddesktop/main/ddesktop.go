package main

import (
	"log"
	"ddesktop/server"
	"net/http"
)

func main() {
	//Start server
	log.Printf("Starting server on http://0.0.0.0:9000...")
	r := server.NewRouter()
	http.Handle("/static/", server.HTTPProvider(http.StripPrefix("/static/", http.FileServer(http.Dir("webroot/static/"))), "GetStatic"))
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":9000", nil))
}
