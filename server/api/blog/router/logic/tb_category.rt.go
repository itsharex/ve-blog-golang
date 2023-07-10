package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type CategoryRouter struct {
	svcCtx *svc.RouterContext
}

func NewCategoryRouter(svcCtx *svc.RouterContext) *CategoryRouter {
	return &CategoryRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Category 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *CategoryRouter) InitCategoryRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.CategoryController
	{
		publicRouter.POST("category/create", handler.CreateCategory)   // 新建Category
		publicRouter.PUT("category/update", handler.UpdateCategory)    // 更新Category
		publicRouter.DELETE("category/delete", handler.DeleteCategory) // 删除Category
		publicRouter.POST("category/find", handler.FindCategory)       // 查询Category

		publicRouter.DELETE("category/deleteByIds", handler.DeleteCategoryByIds) // 批量删除Category列表
		publicRouter.POST("category/list", handler.FindCategoryList)             // 分页查询Category列表
	}
}
