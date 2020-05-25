package db

import (
	"database/sql"
	"fmt"
	"work/book-library/lib/pkg/db/models"

	_ "github.com/lib/pq"
)

type PostgresClient struct {
	Client *sql.DB
}

func (p PostgresClient) Add(dao DAO) error {
	return dao.SQLAdd(p.Client)
}

func (p PostgresClient) Modify(dao DAO, params map[string]interface{}) error {
	return dao.SQLModify(p.Client, params)
}
func (p PostgresClient) Delete(dao DAO) error {
	return dao.SQLDelete(p.Client)
}

func (p PostgresClient) GetAll(dao DAO) ([]interface{}, error) {
	return dao.SQLGetAll(p.Client)
}

func (p PostgresClient) Get(dao DAO, params map[string][]string) ([]interface{}, error) {
	return dao.SQLGet(p.Client, params)
}

func GetPostgresClient() (PostgresClient, func() error, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		models.POSTGRES_HOST, models.POSTGRES_PORT, models.POSTGRES_USER, models.POSTGRES_PASSWORD, models.POSTGRES_DBNAME)

	cli, err := sql.Open(models.POSTGRES_SQL, psqlInfo)
	if err != nil {
		return PostgresClient{}, func() error { return nil }, err
	}
	return PostgresClient{cli}, cli.Close, nil
}
