package dbs

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	DB = sqlx.MustConnect("mysql", "notes:notes@tcp(mysql_db:3306)/weixin")
}

func DBCli() *sqlx.DB {
	return DB
}
