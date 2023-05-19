package main

import (
	"flag"
	pleaco "github.com/pleaco/pleaco"
)

func main() {
	apiAddr := flag.String("api-address", ":8080", "The address to listen on for API HTTP requests.")
	//debug := flag.Bool("debug", false, "enable debug logging")

	flag.Parse()

	go pleaco.runContainers()
	router := pleaco.setupRouter()
	router.Run(*apiAddr)
}
