package wsproxy

import (
	"github.com/spf13/viper"
	"net/http"
	"net"
	"time"
	"log"
	"io"
	"ddesktop/dockerhandler"
	"strconv"
)

var (
	conn net.Conn
	err error
)


func WsProxy() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	log.Println("WebSocket connection opened.")

    	//Start new container
    	containerId := dockerhandler.StartContainer()

    	//Get websockify target
    	target := dockerhandler.GetIP(containerId) + ":" + viper.GetString("container.wsport")

    	//Check if port is open
    	for i := 0; i < 10; i++ {
    		conn, err = net.Dial("tcp", target)
    		if err != nil {
    			if i == 9 {
    				http.Error(w, "Error contacting backend server.", 500)
					log.Printf("Error dialing websocket backend %s: %v", target, err)
					return
    			}
    			time.Sleep(1000 * time.Millisecond)
    		} else{
    			time.Sleep(1000 * time.Millisecond)
    			log.Println("Connected to container " + containerId + " after " + strconv.Itoa(i + 1) + " seconds.")
				break
    		}
    	}

		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "Not a hijacker?", 500)
			return
		}
		nc, _, err := hj.Hijack()
		if err != nil {
			log.Printf("Hijack error: %v", err)
			return
		}
		defer nc.Close()
		defer conn.Close()

		err = r.Write(conn)
		if err != nil {
			log.Printf("Error copying request to target: %v", err)
			return
		}

		errc := make(chan error, 2)
		cp := func(dst io.Writer, src io.Reader) {
			_, err := io.Copy(dst, src)
			errc <- err
		}
		go cp(conn, nc)
		go cp(nc, conn)
		<-errc
		log.Println("WebSocket connection closed.")
		dockerhandler.DeleteContainer(containerId)
	})
}