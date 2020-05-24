package handler

import (
	"fmt"
	"net/http"
)

func deleteBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "calling deleteBooks")
}
