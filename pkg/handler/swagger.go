package handler

import (
	"fmt"
	"net/http"
)

func GetSwagger(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}
