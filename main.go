package main

import (
	"work/book-library/lib/pkg/api"
	"work/book-library/pkg/handler"

	log "github.com/sirupsen/logrus"
)

func main() {
	server := api.Server{
		HostName: "127.0.0.1",
		Port:     "8000",
	}
	err := api.GetMuxRouter(&server)
	if err != nil {
		log.Fatalf("%s", "Error in creating mux router :%v\n", err)
		return
	}
	handler.RegisterRouteHandler(server.MuxRouter)
	err = api.StartAPIServer(server)
	if err != nil {
		log.Fatalf("%s", "Error in main :%v\n", err)
	}
}
