package global

import (
	"github.com/ve-weiyi/go-sdk/utils/glog"
	"github.com/ve-weiyi/ve-admin-store/server/config"
	"github.com/ve-weiyi/ve-admin-store/server/infra/jjwt"
	"github.com/ve-weiyi/ve-admin-store/server/infra/rbac"
	"github.com/ve-weiyi/ve-admin-store/server/utils/timer"

	"github.com/orca-zhang/ecache"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"runtime"
	"strings"
	"sync"
)

var (
	DB     *gorm.DB
	DBList map[string]*gorm.DB
	REDIS  *redis.Client
	CONFIG config.Config
	VP     *viper.Viper
	JWT    *jjwt.JwtToken
	LOG    *glog.Glogger
	Timer  timer.Timer = timer.NewTimerTask()
	//Concurrency_Control             = &singleflight.Group{}

	BlackCache *ecache.Cache
	lock       sync.RWMutex

	//RBAC角色访问控制器
	RbacEnforcer *rbac.CachedEnforcer
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则 使用默认db panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		return DB
		//panic("db no init")
	}
	return db
}

func GetRuntimeRoot() string {
	//获得当前方法运行文件名
	_, filename, _, _ := runtime.Caller(1)
	//找到 resource/language目录
	src := "server"
	index := strings.Index(filename, src)
	root := filename[:index]
	return root
}
