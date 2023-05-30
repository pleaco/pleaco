package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	pleaco "pleaco/pkg"
)

func main() {
	apiAddr := flag.String("api-address", ":8080", "The address to listen on for API HTTP requests.")
	//debug := flag.Bool("debug", false, "enable debug logging")

	flag.Parse()

	go pleaco.RunContainers()
	go pleaco.DeleteContainers()

	router := pleaco.SetupRouter()
	err := router.Run(*apiAddr)
	if err != nil {
		log.Fatal(err)
	}
}
