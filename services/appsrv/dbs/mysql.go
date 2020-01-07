package dbs

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/vnotes/workweixin/services/appsrv/conf"
)

var DB *sqlx.DB

func InitMySQL() {
	DB = sqlx.MustConnect("mysql", fmt.Sprintf("notes:notes@tcp(%s:3306)/weixin", conf.Conf.DBNetWork))
}

func Cli() *sqlx.DB {
	return DB
}
