package initialize

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/ve-weiyi/ve-admin-store/server/global"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	address := redisCfg.Host + ":" + redisCfg.Port
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("Redis 连接失败, err:%v", err)
		return
	}
	client.Set(context.Background(), "connect", time.Now().String(), -1)
	global.REDIS = client

	log.Printf("Redis 连接成功%v! address:%v db:%v", pong, address, redisCfg.DB)
}
