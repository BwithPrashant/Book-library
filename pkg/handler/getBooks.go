package handler

import (
	"fmt"
	"net/http"
	"work/book-library/lib/pkg/api"
	"work/book-library/lib/pkg/db"
	"work/book-library/lib/pkg/db/dao"
	"work/book-library/lib/pkg/db/models"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	book := dao.Book{}

	//Get postgres client
	dbClient, cleanfunc, err := db.GetClient(models.POSTGRES_SQL)
	defer cleanfunc()
	if err != nil {
		log.Errorf("Failed fetching db client. Error:%s\n", err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Failed fetching db client. Error:%s", err.Error()), nil, w, 400)
		return
	}
	resp, err := dbClient.GetAll(book)
	if err != nil {
		log.Errorf("Failed to fetch details of books. Error:%s\n", err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Failed to fetch details of books. Error:%s\n", err.Error()), nil, w, 400)
		return
	}
	api.WriteRestResponse(api.SUCCESS, "Book data successfully retrieved", resp, w, 200)
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
