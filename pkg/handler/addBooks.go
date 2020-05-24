package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"work/book-library/lib/pkg/api"
	"work/book-library/pkg/objects"

	log "github.com/sirupsen/logrus"
)

func ValidateAddBooksRequest(addBookRequest objects.AddBooksRequest) error {
	if addBookRequest.Author == "" {
		return fmt.Errorf("author name can't be empty")
	}

	if addBookRequest.ISBN == "" {
		return fmt.Errorf("ISBN can't be empty")
	}

	if addBookRequest.Title == "" {
		return fmt.Errorf("Title can't be empty")
	}
	return nil
}

func addBooks(w http.ResponseWriter, r *http.Request) {
	var addBookRequest objects.AddBooksRequest
	err := json.NewDecoder(r.Body).Decode(&addBookRequest)
	if err != nil {
		log.Errorf("Error in parsing request. Error:%v\n", err)
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Error in parsing request. Error:%v", err), nil, w, 400)
		return
	}
	err = ValidateAddBooksRequest(addBookRequest)
	if err != nil {
		log.Errorf("Error in Validating addBooks request. Error:%v\n", err)
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Error in Validating addBooks request. Error:%v", err), nil, w, 400)
		return
	}
	log.Infof("Request is : %v\n", addBookRequest)
	//TODO:- Save request to DB
	api.WriteRestResponse(api.SUCCESS, "Book entry successfully saved", addBookRequest, w, 202)
}
