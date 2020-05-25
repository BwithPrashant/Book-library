package handler

import (
	"fmt"
	"net/http"
	"strings"
	"work/book-library/lib/pkg/api"
	"work/book-library/lib/pkg/db"
	"work/book-library/lib/pkg/db/dao"
	"work/book-library/lib/pkg/db/models"
	"work/book-library/lib/pkg/utils"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	supportedFilters := []string{"title", "author", "country", "page"}

	var message string
	var queryParams map[string][]string
	queryParams = make(map[string][]string)
	for queryStringKey, sqlQueryKey := range params {
		if !utils.IsStringInSlice(strings.ToLower(queryStringKey), supportedFilters) {
			message = fmt.Sprintf("Unsupported filter: %v", queryStringKey)
			api.WriteRestResponse(api.ERROR, message, nil, w, http.StatusBadRequest)
			return
		}
		queryParams[queryStringKey] = sqlQueryKey
	}

	client, cleanfunc, err := db.GetClient(models.POSTGRES_SQL)
	defer cleanfunc()
	if err != nil {
		log.Errorf("Failed fetching db client. Error:%s\n", err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Failed fetching db client. Error:%s", err.Error()), nil, w, http.StatusInternalServerError)
		return
	}

	book := dao.Book{}
	fetchedBooks, err := client.Get(book, queryParams)
	if err != nil {
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("failed fetching job details. Error: %s", err.Error()), nil, w, http.StatusInternalServerError)
		return
	}

	if fetchedBooks == nil {
		api.WriteRestResponse(api.ERROR, "Details Not Found", nil, w, http.StatusNotFound)
		return
	}

	api.WriteRestResponse(api.SUCCESS, "Booklist successfully fetched", fetchedBooks, w, http.StatusOK)
	return
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	bookId, err := uuid.Parse(urlParams["id"])
	if err != nil {
		log.Errorf("Invalid request id %s. Error:%s\n", urlParams["id"], err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Invalid request id %s. Error:%s\n", urlParams["id"], err.Error()), nil, w, 400)
		return
	}
	book := dao.Book{}

	//Get postgres client
	dbClient, cleanfunc, err := db.GetClient(models.POSTGRES_SQL)
	defer cleanfunc()
	if err != nil {
		log.Errorf("Failed fetching db client. Error:%s\n", err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Failed fetching db client. Error:%s", err.Error()), nil, w, 400)
		return
	}
	resp, err := dbClient.Get(book, map[string][]string{
		"Id": {
			bookId.String(),
		},
	})
	if err != nil {
		log.Errorf("Failed to find details for book with id %s. Error:%s\n", bookId, err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Failed to find details for book with id %s. Error:%s", bookId, err.Error()), nil, w, 400)
		return
	}
	api.WriteRestResponse(api.SUCCESS, "Book data successfully retrieved", resp, w, 200)
}
