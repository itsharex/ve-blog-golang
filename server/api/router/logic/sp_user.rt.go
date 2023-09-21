package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/router/svc"
)

type UserRouter struct {
	svcCtx *svc.RouterContext
}

func NewUserRouter(ctx *svc.RouterContext) *UserRouter {
	return &UserRouter{
		svcCtx: ctx,
	}
}

// InitUserRouter 初始化 User 路由信息
func (s *UserRouter) InitUserRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.UserController
	{
		loginRouter.GET("user/login_history", handler.FindUserLoginHistory) // 用户信息
		loginRouter.GET("user/info", handler.GetUserInfo)                   // 用户信息
		loginRouter.GET("user/menus", handler.GetUserMenus)                 // 用户菜单
		loginRouter.GET("user/apis", handler.GetUserApis)                   // 用户资源

		loginRouter.POST("user/info", handler.UpdateUserInfo)            // 更新用户信息
		loginRouter.POST("user/avatar", handler.UpdateUserAvatar)        // 更新用户头像
		loginRouter.POST("user/update_status", handler.UpdateUserStatus) // 更新用户状态

		// 管理员操作
		loginRouter.POST("user/list", handler.FindUserList)            // 获取用户列表
		loginRouter.POST("user/list/areas", handler.FindUserListAreas) // 获取用户地区
	}
}