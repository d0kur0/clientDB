package redisWrapper

import (
	"context"
	"log"

	"github.com/d0kur0/clientDB/utils/configMgr"

	"github.com/pkg/errors"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func Init() error {
	config := configMgr.Get()

	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password,
		DB:       config.Redis.Dbname,
	})

	_, err := rdb.Ping(ctx).Result()

	if err == nil {
		log.Println("Connected to redis successful")
	}

	return errors.Wrap(err, "Error with connect to redis")
}
