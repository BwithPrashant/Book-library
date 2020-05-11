package main

import (
	"fmt"
	"work/book-library/lib/api"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello")
	server := api.Server{
		HostName: "127.0.0.1",
		Port:     "8000",
	}
	err := api.StartAPIServer(server)
	if err != nil {
		log.Fatalf("%s", "Error in main :%v\n", err)
	}

}
