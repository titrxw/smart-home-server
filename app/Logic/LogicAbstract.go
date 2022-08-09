package logic

import (
	"github.com/go-redis/redis/v8"
	global "github.com/titrxw/go-framework/src/Global"
	helper "github.com/titrxw/smart-home-server/app/Helper"
	"gorm.io/gorm"
	"time"
)

type LogicAbstract struct {
}

func (logic LogicAbstract) GetDefaultDb() *gorm.DB {
	return global.FApp.DbFactory.Channel("default")
}

func (logic LogicAbstract) GetOperateOrReportNumber(appId string) string {
	return helper.Sha1(appId + helper.UUid() + time.Now().String())
}

func (logic LogicAbstract) GetDefaultRedis() *redis.Client {
	return global.FApp.RedisFactory.Channel("default")
}
