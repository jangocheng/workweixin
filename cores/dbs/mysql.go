package dbs

import "github.com/jmoiron/sqlx"

var DB *sqlx.DB


func init() {
	DB = sqlx.MustConnect("mysql", "")
}
