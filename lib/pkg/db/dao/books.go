package dao

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"work/book-library/lib/pkg/db/models"
)

const (
	DEFAULT_RECORDS_LIMIT = 20
	DEFAULT_PAGE_NUM      = "1"
)

type Book struct {
	Id      string
	Isbn    string
	Title   string
	Author  string
	Country string
}

func (b Book) SQLAdd(cli *sql.DB) error {
	query := "INSERT INTO book_library.books (id,isbn,title,author,country) VALUES ($1, $2, $3, $4, $5);"
	stmt, err := cli.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(b.Id, b.Isbn, b.Title, b.Author, b.Country)
	return err
}

func (b Book) SQLModify(cli *sql.DB, params map[string]interface{}) error {
	if len(params) == 0 {
		return nil
	}
	idx := 1
	query := "UPDATE book_library.books SET "
	var args []interface{}
	for k, v := range params {
		query += fmt.Sprintf("%s=$%s,", k, strconv.Itoa(idx))
		args = append(args, v)
		idx++
	}
	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE id='%s';", b.Id)

	fmt.Println("query is : ", query)
	stmt, err := cli.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("Data not found in database")
	}
	return nil
}

func (b Book) SQLDelete(cli *sql.DB) error {
	query := "DELETE FROM book_library.books WHERE id = $1;"
	stmt, err := cli.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(b.Id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("Data not found in database")
	}
	return nil
}

func (b Book) SQLGetAll(cli *sql.DB) ([]interface{}, error) {
	books := []interface{}{}
	var bookResp models.BookResponse
	query := "SELECT * FROM book_library.books;"
	stmt, err := cli.Prepare(query)
	if err != nil {
		return books, err
	}

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		return books, err
	}

	for rows.Next() {
		err := rows.Scan(&bookResp.Id, &bookResp.Isbn, &bookResp.Title, &bookResp.Author, &bookResp.Country)
		if err != nil {
			return books, err
		}
		book := Book{bookResp.Id.String, bookResp.Isbn.String, bookResp.Title.String, bookResp.Author.String, bookResp.Country.String}
		books = append(books, book)
	}
	return books, nil
}

//applyOffset returns offset as per page number
func applyOffset(pg int) int {
	if pg <= 0 {
		return 0
	}
	offset := (pg - 1) * DEFAULT_RECORDS_LIMIT
	return offset
}

func (b Book) SQLGet(cli *sql.DB, params map[string][]string) ([]interface{}, error) {
	var books []interface{}
	var bookResp models.BookResponse

	pageNo := DEFAULT_PAGE_NUM
	pageNoParam, ok := params["page"]
	if ok && len(pageNoParam) >= 0 {
		pageNo = pageNoParam[0]
	}
	delete(params, "page")

	query := "SELECT * FROM book_library.books"
	if len(params) != 0 {
		query += " WHERE "
	}
	idx := 1
	var args []interface{}
	for filterName, filterParamSlice := range params {
		for _, filterParam := range filterParamSlice {
			var delimeter string = ","
			filterSlice := strings.Split(filterParam, delimeter)
			query += fmt.Sprintf("%s IN (", filterName)

			for _, filterVal := range filterSlice {
				args = append(args, filterVal)
				query += fmt.Sprintf("$%s,", strconv.Itoa(idx))
				idx++
			}
			query = query[:len(query)-1]
			query += fmt.Sprintf(") AND ")
		}
	}

	pg, err := strconv.Atoi(pageNo)
	if err != nil {
		return books, err
	}

	if len(params) != 0 {
		query = query[:len(query)-4]
	}
	query += " limit " + fmt.Sprintf("%d offset %d;", DEFAULT_RECORDS_LIMIT, applyOffset(pg))

	stmt, err := cli.Prepare(query)
	if err != nil {
		return books, err
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&bookResp.Id, &bookResp.Isbn, &bookResp.Title, &bookResp.Author, &bookResp.Country)
		if err != nil {
			return books, err
		}
		book := Book{bookResp.Id.String, bookResp.Isbn.String, bookResp.Title.String, bookResp.Author.String, bookResp.Country.String}
		books = append(books, book)
	}
	err = rows.Err()
	if err != nil {
		return books, err
	}
	if len(books) == 0 {
		return books, fmt.Errorf("Data not found")
	}
	return books, nil
}
