package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"work/book-library/lib/pkg/api"
	"work/book-library/lib/pkg/db"
	"work/book-library/lib/pkg/db/dao"
	"work/book-library/lib/pkg/db/models"
	"work/book-library/pkg/objects"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

func ValidateAddBooksRequest(addBookRequest objects.AddBookRequest) error {
	if addBookRequest.Author == "" {
		return fmt.Errorf("author name can't be empty")
	}

	if addBookRequest.Isbn == "" {
		return fmt.Errorf("ISBN can't be empty")
	}

	if addBookRequest.Title == "" {
		return fmt.Errorf("Title can't be empty")
	}
	return nil
}

func AddBooks(w http.ResponseWriter, r *http.Request) {
	var addBookRequest objects.AddBookRequest
	//Decode request body
	err := json.NewDecoder(r.Body).Decode(&addBookRequest)
	if err != nil {
		log.Errorf("Failed to parse request body. Error:%s\n", err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Failed to parse request body. Error:%s", err.Error()), nil, w, 400)
		return
	}

	log.Infof("Add Book request is : %v\n", addBookRequest)

	//Validate request
	err = ValidateAddBooksRequest(addBookRequest)
	if err != nil {
		log.Errorf("Error in Validating addBooks request. Error:%s\n", err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Error in Validating addBooks request. Error:%s", err.Error()), nil, w, 400)
		return
	}

	//Get postgres client
	dbClient, cleanfunc, err := db.GetClient(models.POSTGRES_SQL)
	defer cleanfunc()
	if err != nil {
		log.Errorf("Failed fetching db client. Error:%s\n", err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Failed fetching db client. Error:%s", err.Error()), nil, w, 400)
		return
	}

	bookId := uuid.New().String()

	//Save Data to DB
	book := dao.Book{bookId, addBookRequest.Isbn, addBookRequest.Title, addBookRequest.Author, addBookRequest.Country}
	err = dbClient.Add(book)
	if err != nil {
		log.Errorf("Failed in saving data to DB. Error:%s\n", err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Failed in saving data to DB. Error:%s", err.Error()), nil, w, 400)
		return
	}

	//Response to be returned to end user
	bookIdentity := objects.BookIdentity{
		Id: bookId,
		Data: objects.AddBookRequest{
			Isbn:    addBookRequest.Isbn,
			Title:   addBookRequest.Title,
			Author:  addBookRequest.Author,
			Country: addBookRequest.Country,
		},
	}
	api.WriteRestResponse(api.SUCCESS, "Book entry successfully saved", bookIdentity, w, 202)
}
