package core

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 ve77
	当前版本:v2.5.5
    加群方式:微信号：shouzi_1994 QQ群：622360840
	插件市场:https://plugin.ve77.com
	GVA讨论社区:https://support.qq.com/products/371961
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:9090
	如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.ve77.com/coffee/index.html
`, address)
	global.LOG.Error(s.ListenAndServe().Error())
}
