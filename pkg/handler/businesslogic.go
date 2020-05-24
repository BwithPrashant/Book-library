package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	GetAPI = map[string]map[string]http.HandlerFunc{
		"/books": map[string]http.HandlerFunc{
			http.MethodGet: getBooks,
		},
		"/books/{id}": map[string]http.HandlerFunc{
			http.MethodGet: getBookByID,
		},
		"/swagger": map[string]http.HandlerFunc{
			http.MethodGet: getSwagger,
		},
		"/health": map[string]http.HandlerFunc{
			http.MethodGet: health,
		},
	}

	PostAPI = map[string]map[string]http.HandlerFunc{
		"/books": map[string]http.HandlerFunc{
			http.MethodPost: addBooks,
		},
	}

	PutAPI = map[string]map[string]http.HandlerFunc{
		"/books": map[string]http.HandlerFunc{
			http.MethodPut: updateBooks,
		},
	}

	DeleteAPI = map[string]map[string]http.HandlerFunc{
		"/books": map[string]http.HandlerFunc{
			http.MethodDelete: deleteBooks,
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
