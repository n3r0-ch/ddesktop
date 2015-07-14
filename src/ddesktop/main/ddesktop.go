package main

import (
	"log"
	"github.com/spf13/viper"
	"ddesktop/server"
	"ddesktop/wsproxy"
	"ddesktop/dockerhandler"
	"net/http"
	auth "github.com/abbot/go-http-auth"
	"os"
)

func main() {

	pwd, _ := os.Getwd()

	//Set config path and type
	viper.SetConfigName("config")
	viper.AddConfigPath(pwd)
	viper.AddConfigPath("/etc/ddesktop/")
	viper.SetConfigType("yaml")

	//Read config
	err := viper.ReadInConfig()
	if err != nil { 
	    log.Fatalln(err)
	}

	log.Println(viper.GetString("container.prefix"))

	//Cleanup existing containers
	dockerhandler.CleanUp()

	//Pull new docker image
	if viper.GetBool("container.pull"){
		dockerhandler.PullImage()
	}


	//Get authentication setting
	htpasswd := auth.HtpasswdFileProvider(viper.GetString("htpasswd.path"))
	authenticator := auth.NewBasicAuthenticator(".ddesktop", htpasswd)

	//Start server
	log.Printf("Starting server on http://0.0.0.0:" + viper.GetString("server.port.http") + " and https://0.0.0.0:" + viper.GetString("server.port.https") + "...")
	http.Handle("/websockify", auth.JustCheck(authenticator, wsproxy.WsProxy()))
	http.HandleFunc("/", auth.JustCheck(authenticator, server.Static()))

	go func() {
		if err := http.ListenAndServeTLS(":" + viper.GetString("server.port.https"), viper.GetString("ssl.cert"), viper.GetString("ssl.key"), nil); err != nil {
			log.Fatalln(err)
		}
	}();
	if err := http.ListenAndServe(":" + viper.GetString("server.port.http"), http.HandlerFunc(server.RedirectHttps)); err != nil {
		log.Fatalln(err)
	}
}
