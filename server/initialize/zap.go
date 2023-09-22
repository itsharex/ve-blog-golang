package initialize

import (
	"fmt"
	"os"

	"github.com/ve-weiyi/ve-blog-golang/server/global"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/copy"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/files"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/glog"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/glog/zaplog"
)

// Zap 获取 zap.Logger
func Zap() {
	if ok, _ := files.PathExists(global.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}

	cfg := zaplog.ZapConfig{}

	copy.DeepCopyByJson(global.CONFIG.Zap, &cfg)

	glog.ReplaceZapGlobals(cfg)
	global.LOG = glog.NewGlogger(1, cfg)

	global.LOG.Printf("日志组件初始化成功！")
	return
}
