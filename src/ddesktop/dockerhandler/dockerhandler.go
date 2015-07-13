package dockerhandler

import (
	"github.com/samalba/dockerclient"
	"github.com/spf13/viper"
	"log"
	"strings"
	gouuid "github.com/nu7hatch/gouuid"
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
	    		DeleteContainer(c.Id)
	    	}
    	}
    }
}

func StartContainer() string{
	client := GetClient()

	//Get uuid
	uuid, err := gouuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}

	//Get name
	name := viper.GetString("container.prefix") + uuid.String()

	//Create container
	containerConfig := &dockerclient.ContainerConfig{
        Image: viper.GetString("container.image"),
    }
    containerId, err := client.CreateContainer(containerConfig, name)
    if err != nil {
        log.Println(err)
    }

    //Start container
    hostConfig := &dockerclient.HostConfig{}
    err = client.StartContainer(containerId, hostConfig)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Started container with id " + containerId)

    return containerId
}

func GetIP(containerId string) string{
	client := GetClient()
	info, err := client.InspectContainer(containerId)
	if err != nil {
        log.Println(err)
    }
	return info.NetworkSettings.IPAddress
}

func DeleteContainer(containerId string){
	log.Println("Removing container " + containerId)
	client := GetClient()
	client.KillContainer(containerId, "SIGKILL")
	client.RemoveContainer(containerId, true, false)
}