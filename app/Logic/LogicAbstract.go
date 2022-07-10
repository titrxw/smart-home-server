package logic

import (
	"github.com/go-redis/redis/v8"
	global "github.com/titrxw/go-framework/src/Global"
	"gorm.io/gorm"
)

type LogicAbstract struct {
}

func (logic LogicAbstract) GetDefaultDb() *gorm.DB {
	return global.FApp.DbFactory.Channel("default")
}

func (logic LogicAbstract) GetDefaultRedis() *redis.Client {
	return global.FApp.RedisFactory.Channel("default")
}
