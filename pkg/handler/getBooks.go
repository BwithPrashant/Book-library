package handler

import (
	"fmt"
	"net/http"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "calling getBooks")
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "calling getBookByID")
}
