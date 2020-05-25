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
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func ValidateModifyBooksRequest(modifyBookRequest objects.ModifyBookRequest) error {
	if modifyBookRequest.Author == "" {
		return fmt.Errorf("author name can't be empty")
	}

	if modifyBookRequest.Title == "" {
		return fmt.Errorf("Title can't be empty")
	}
	return nil
}

func SetModifyRequestParams(modifyBookRequest objects.ModifyBookRequest) map[string]interface{} {
	params := make(map[string]interface{})
	if modifyBookRequest.Author != "" {
		params["Author"] = modifyBookRequest.Author
	}

	if modifyBookRequest.Title != "" {
		params["Title"] = modifyBookRequest.Title
	}

	if modifyBookRequest.Country != "" {
		params["Country"] = modifyBookRequest.Country
	}

	return params
}
func UpdateBooks(w http.ResponseWriter, r *http.Request) {

	urlParams := mux.Vars(r)
	bookId, err := uuid.Parse(urlParams["id"])
	if err != nil {
		log.Errorf("Invalid request id %s. Error:%s\n", urlParams["id"], err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Invalid request id %s. Error:%s\n", urlParams["id"], err.Error()), nil, w, 400)
		return
	}
	book := dao.Book{Id: bookId.String()}

	var modifyBookRequest objects.ModifyBookRequest
	//Decode request body
	err = json.NewDecoder(r.Body).Decode(&modifyBookRequest)
	if err != nil {
		log.Errorf("Failed to parse request body. Error:%s\n", err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Failed to parse request body. Error:%s", err.Error()), nil, w, 400)
		return
	}

	log.Infof("Modify book request is : %v\n", modifyBookRequest)

	//Validate request
	err = ValidateModifyBooksRequest(modifyBookRequest)
	if err != nil {
		log.Errorf("Error in Validating addBooks request. Error:%s\n", err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Error in Validating addBooks request. Error:%s", err.Error()), nil, w, 400)
		return
	}

	params := SetModifyRequestParams(modifyBookRequest)

	//Get postgres client
	dbClient, cleanfunc, err := db.GetClient(models.POSTGRES_SQL)
	defer cleanfunc()
	if err != nil {
		log.Errorf("Failed fetching db client. Error:%s\n", err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Failed fetching db client. Error:%s", err.Error()), nil, w, 400)
		return
	}
	err = dbClient.Modify(book, params)
	if err != nil {
		log.Errorf("Failed to find details for book with id %s. Error:%s\n", bookId, err.Error())
		api.WriteRestResponse(api.ERROR, fmt.Sprintf("Failed to find details for book with id %s. Error:%s", bookId, err.Error()), nil, w, 400)
		return
	}
	api.WriteRestResponse(api.SUCCESS, "Book data successfully modified", nil, w, 200)
}
