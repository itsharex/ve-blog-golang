package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/router/svc"
)

type BlogRouter struct {
	svcCtx *svc.RouterContext
}

func NewBlogRouter(ctx *svc.RouterContext) *BlogRouter {
	return &BlogRouter{
		svcCtx: ctx,
	}
}

// 初始化 Blog 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *BlogRouter) InitBlogRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.BlogController
	{
		publicRouter.GET("ws", handler.WebSocket)                  // websocket
		publicRouter.GET("about", handler.GetAboutMe)              // 查询关于我
		publicRouter.POST("chat/records", handler.FindChatRecords) // 查询聊天记录

		loginRouter.GET("home", handler.GetHomeInfo)                   // 获取首页信息
		loginRouter.POST("about", handler.UpdateAboutMe)               // 更新关于我
		loginRouter.GET("admin/website/config", handler.GetConfig)     // 获取网站配置
		loginRouter.POST("admin/website/config", handler.UpdateConfig) // 更新网站配置
	}
}