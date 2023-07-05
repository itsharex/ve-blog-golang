package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/go-sdk/utils/glog"
	"github.com/ve-weiyi/ve-admin-store/server/config"
	"github.com/ve-weiyi/ve-admin-store/server/global"
	"gorm.io/gorm"
)

// 注册需要用到的gorm、redis、model
type RepositoryContext struct {
	Config  *config.Config
	DbEngin *gorm.DB
	DBList  map[string]*gorm.DB
	Cache   *redis.Client
	Log     *glog.Glogger
	//下面是一些Model
}

func NewRepositoryContext(cfg *config.Config) *RepositoryContext {
	return &RepositoryContext{
		Config:  cfg,
		DbEngin: global.DB,
		DBList:  global.DBList,
		Cache:   global.REDIS,
		Log:     global.LOG,
	}
}
