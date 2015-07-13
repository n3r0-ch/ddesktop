package main

import (
	"log"
	"github.com/spf13/viper"
	"ddesktop/server"
	"ddesktop/wsproxy"
	"ddesktop/dockerhandler"
	"net/http"
)

func main() {

	//Set config path and type
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/ddesktop/")
	viper.SetConfigType("yaml")

	//Set config defaults
	viper.SetDefault("server.port", "9000")
	viper.SetDefault("container.image", "n3r0ch/ddesktop-min")
	viper.SetDefault("container.prefix", "dd--")

	//Read config
	err := viper.ReadInConfig()
	if err != nil { 
	    log.Fatalln(err)
	}

	//Cleanup existing containers
	dockerhandler.CleanUp()

	//Pull new docker image
	dockerhandler.PullImage()

	//Start server
	log.Printf("Starting server on http://0.0.0.0:" + viper.GetString("server.port") + "...")
	r := server.NewRouter()
	http.Handle("/static/", server.HTTPProvider(http.StripPrefix("/static/", http.FileServer(http.Dir("webroot/static/"))), "GetStatic"))
	http.Handle("/websockify", wsproxy.WsProxy())
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":" + viper.GetString("server.port"), nil))
}
