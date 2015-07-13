package main

import (
	"log"
	"github.com/spf13/viper"
	"ddesktop/server"
	"ddesktop/wsproxy"
	"net/http"
)

func main() {

	//Set config path and type
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/ddesktop/")
	viper.SetConfigType("yaml")

	//Set config defaults
	viper.SetDefault("server.port", "9000")
	viper.SetDefault("n3r0ch/ddesktop-min", "9000")

	//Read config
	err := viper.ReadInConfig()
	if err != nil { 
	    log.Fatalln(err)
	}

	//Start server
	log.Printf("Starting server on http://0.0.0.0:" + viper.GetString("server.port") + "...")
	r := server.NewRouter()
	http.Handle("/static/", server.HTTPProvider(http.StripPrefix("/static/", http.FileServer(http.Dir("webroot/static/"))), "GetStatic"))
	http.Handle("/websockify", wsproxy.WsProxy("localhost:6080"))
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":" + viper.GetString("server.port"), nil))
}
