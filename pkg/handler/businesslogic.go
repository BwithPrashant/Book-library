package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	GetAPI = map[string]map[string]http.HandlerFunc{
		"/books": map[string]http.HandlerFunc{
			http.MethodGet: GetBooks,
		},
		"/books/{id}": map[string]http.HandlerFunc{
			http.MethodGet: GetBookByID,
		},
		"/swagger": map[string]http.HandlerFunc{
			http.MethodGet: GetSwagger,
		},
		"/health": map[string]http.HandlerFunc{
			http.MethodGet: Health,
		},
	}

	PostAPI = map[string]map[string]http.HandlerFunc{
		"/books": map[string]http.HandlerFunc{
			http.MethodPost: AddBooks,
		},
	}

	PutAPI = map[string]map[string]http.HandlerFunc{
		"/books/{id}": map[string]http.HandlerFunc{
			http.MethodPut: UpdateBooks,
		},
	}

	DeleteAPI = map[string]map[string]http.HandlerFunc{
		"/books/{id}": map[string]http.HandlerFunc{
			http.MethodDelete: DeleteBooks,
		},
	}
)

func RegisterRouteHandler(router *mux.Router) {

	for route, routeHandlerMap := range GetAPI {
		for verb, handlerFunc := range routeHandlerMap {
			router.HandleFunc(route, handlerFunc).Methods(verb)
		}
	}

	for route, routeHandlerMap := range PostAPI {
		for verb, handlerFunc := range routeHandlerMap {
			router.HandleFunc(route, handlerFunc).Methods(verb)
		}
	}

	for route, routeHandlerMap := range PutAPI {
		for verb, handlerFunc := range routeHandlerMap {
			router.HandleFunc(route, handlerFunc).Methods(verb)
		}
	}

	for route, routeHandlerMap := range DeleteAPI {
		for verb, handlerFunc := range routeHandlerMap {
			router.HandleFunc(route, handlerFunc).Methods(verb)
		}
	}
}
