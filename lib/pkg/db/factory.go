package db

import (
	"database/sql"
	"fmt"
	"work/book-library/lib/pkg/db/models"
)

type DAO interface {
	SQLAdd(*sql.DB) error
	SQLModify(*sql.DB, map[string]interface{}) error
	SQLDelete(*sql.DB) error
	SQLGetAll(*sql.DB) ([]interface{}, error)
	SQLGet(*sql.DB, map[string][]string) ([]interface{}, error)
}

type Client interface {
	Add(DAO) error
	Modify(DAO, map[string]interface{}) error
	Delete(DAO) error
	GetAll(DAO) ([]interface{}, error)
	Get(DAO, map[string][]string) ([]interface{}, error)
}

func GetClient(dbName string) (Client, func() error, error) {
	switch dbName {
	case models.POSTGRES_SQL:
		return GetPostgresClient()
	default:
		return nil, nil, fmt.Errorf("db provided is not supported")
	}
}
