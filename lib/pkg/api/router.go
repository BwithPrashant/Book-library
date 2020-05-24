package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func NewRestJSONResponse(status RestStatus, message string, data interface{}) RestJSONResponse {
	resp := RestJSONResponse{
		Status:  string(status),
		Message: message,
	}
	if data != nil {
		resp.Data = data
	}
	return resp
}

func WriteRestResponse(status RestStatus, message string, data interface{}, w http.ResponseWriter, statuscode int) {
	resp := NewRestJSONResponse(status, message, data)
	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Errorf("Error in marshalling json resposne %v\n", err)
		http.Error(w, fmt.Sprintf("failed to marshal resp: %s. Error: %v", resp, err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	w.Write(respBytes)
}

func GetMuxRouter(server *Server) error {
	router := mux.NewRouter()
	if router == nil {
		return fmt.Errorf("Error in creating new mux router instance")
	}
	// Attach the mux router to the route path of the API server
	http.Handle("/", router)
	server.MuxRouter = router
	return nil
}

func StartAPIServer(server Server) error {
	log.Infof("API server is starting at %s:%s\n", server.HostName, server.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", server.HostName, server.Port), nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s", "Error in starting api server. Error : %v\n", err))
		return err
	}
	return nil
}
