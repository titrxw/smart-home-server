package logic

import (
	"github.com/redis/go-redis/v9"
	"github.com/titrxw/smart-home-server/app/common/helper"
	"github.com/we7coreteam/w7-rangine-go-support/src/facade"
	"gorm.io/gorm"
	"time"
)

type Abstract struct {
}

func (logic Abstract) GetDefaultDb() *gorm.DB {
	db, _ := facade.GetDbFactory().Channel("default")
	return db
}

func (logic Abstract) GetOperateOrReportNumber(appId string) string {
	return helper.Sha1(appId + helper.UUid() + time.Now().String())
}

func (logic Abstract) GetDefaultRedis() redis.Cmdable {
	redis, _ := facade.GetRedisFactory().Channel("default")

	return redis
}
