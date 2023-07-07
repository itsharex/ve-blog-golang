package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type AdminRouter struct {
	svcCtx *svc.RouterContext
}

func NewAdminRouter(ctx *svc.RouterContext) *AdminRouter {
	return &AdminRouter{
		svcCtx: ctx,
	}
}

// 初始化 Api 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *AdminRouter) InitAdminRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.AdminController

	{
		loginRouter.GET("roles", handler.GetRoleTreeList)                      // 获取Role列表
		loginRouter.POST("role/update_menus", handler.UpdateRoleMenus)         // 获取Role列表
		loginRouter.POST("role/update_resources", handler.UpdateRoleResources) // 获取Role列表
	}
	{
		loginRouter.GET("apis", handler.GetApiTreeList) // 获取Api列表
	}
}