package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}
func getMuxRouter() (*mux.Router, error) {
	router := mux.NewRouter()
	if router == nil {
		return nil, fmt.Errorf("Error in creating new mux router instance")
	}

	return router, nil
}
func StartAPIServer(server Server) error {
	log.Debug("Starting api server")
	router, err := getMuxRouter()
	// Attach the mux router to the route path of the API server
	http.Handle("/", router)
	router.HandleFunc("/", welcome)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s", "Error in starting api server. Error : %v\n", err))
		return err
	}
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", server.HostName, server.Port), nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s", "Error in starting api server. Error : %v\n", err))
		return err
	}
	return nil
}
