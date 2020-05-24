package api

import "github.com/gorilla/mux"

type Server struct {
	HostName  string
	Port      string
	MuxRouter *mux.Router
}

type RestStatus string

const (
	ERROR   RestStatus = "error"
	SUCCESS RestStatus = "success"
)

type RestJSONResponse struct {
	Status  string      `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}
