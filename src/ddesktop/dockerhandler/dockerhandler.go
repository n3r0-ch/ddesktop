package dockerhandler

import (
	"github.com/samalba/dockerclient"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func GetClient() dockerclient.DockerClient {
	docker, _ := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	return *docker
}

func PullImage() {
	client := GetClient()
	log.Println("Pulling image: " + viper.GetString("container.image") + ". This can take several minutes.")
	err := client.PullImage(viper.GetString("container.image"), nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Finished dowloading image: " + viper.GetString("container.image") + ".")
}

func CleanUp(){
	client := GetClient()
	log.Println("Remove all existing containers...")

	containers, err := client.ListContainers(true, false, "")
	if err != nil {
        log.Fatalln(err)
    }
    for _, c := range containers {
    	if len(c.Names) > 0 {
    		if strings.HasPrefix(c.Names[0], "/" + viper.GetString("container.prefix")) {
	    		log.Println("Remove container " + c.Names[0])
	    		client.KillContainer(c.Id, "SIGKILL")
	    		client.RemoveContainer(c.Id, true, false)
	    	}
    	}
    }

}