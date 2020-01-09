package todos

import (
	"time"

	"github.com/vnotes/workweixin/services/appsrv/dbs"
)

const (
	todoListKey    = "weixin:appsrv:todolist"
	todoListExpire = 1 * time.Hour
)

func getToDoListByCache() (string, error) {
	val, err := dbs.RDBCli().Get(todoListKey).Result()
	if err != nil {
		return "", err
	}
	return val, err
}

func CacheToDoList(data string) error {
	return dbs.RDBCli().Set(todoListKey, data, todoListExpire).Err()
}

func DelToDoList() error {
	return dbs.RDBCli().Del(todoListKey).Err()
}