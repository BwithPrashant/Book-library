package handler

import (
	"fmt"
	"net/http"
)

func updateBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "calling updateBooks")
}
