package dbs

import (
	"fmt"
	"log"

	"github.com/vnotes/workweixin/services/appsrv/conf"

	"github.com/go-redis/redis/v7"
)

var rdb *redis.Client

func InitRedis() {
	addr := fmt.Sprintf("%s:6379", conf.Conf.RedisWork)
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalf("init redis error %#v", err)
	}

}

func RDBCli() *redis.Client {
	return rdb
}
