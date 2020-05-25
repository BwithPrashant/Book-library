package models

import "database/sql"

type BookResponse struct {
	Id      sql.NullString
	Isbn    sql.NullString
	Title   sql.NullString
	Author  sql.NullString
	Country sql.NullString
}
