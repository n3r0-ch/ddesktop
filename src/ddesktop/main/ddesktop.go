package main

import (
	"log"
	"github.com/spf13/viper"
	"ddesktop/server"
	"ddesktop/wsproxy"
	"ddesktop/dockerhandler"
	"net/http"
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
	dockerhandler.PullImage()

	//Start server
	log.Printf("Starting server on http://0.0.0.0:" + viper.GetString("server.port.http") + " and https://0.0.0.0:" + viper.GetString("server.port.https") + "...")
	r := server.NewRouter()
	http.Handle("/static/", server.CacheProvider(server.HTTPProvider(http.StripPrefix("/static/", http.FileServer(http.Dir("webroot/static/"))), "GetStatic")))
	http.Handle("/websockify", wsproxy.WsProxy())
	http.Handle("/", r)


	go func() {
		if err := http.ListenAndServeTLS(":" + viper.GetString("server.port.https"), viper.GetString("ssl.cert"), viper.GetString("ssl.key"), nil); err != nil {
			log.Fatalln(err)
		}
	}();
	if err := http.ListenAndServe(":" + viper.GetString("server.port.http"), http.HandlerFunc(server.RedirectHttps)); err != nil {
		log.Fatalln(err)
	}
}
