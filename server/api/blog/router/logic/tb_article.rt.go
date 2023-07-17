package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type ArticleRouter struct {
	svcCtx *svc.RouterContext
}

func NewArticleRouter(svcCtx *svc.RouterContext) *ArticleRouter {
	return &ArticleRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Article 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *ArticleRouter) InitArticleRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.ArticleController
	{
		publicRouter.POST("article/create", handler.CreateArticle)   // 新建Article
		publicRouter.PUT("article/update", handler.UpdateArticle)    // 更新Article
		publicRouter.DELETE("article/delete", handler.DeleteArticle) // 删除Article
		publicRouter.POST("article/find", handler.FindArticle)       // 查询Article

		publicRouter.DELETE("article/deleteByIds", handler.DeleteArticleByIds) // 批量删除Article列表
		publicRouter.POST("article/list", handler.FindArticleList)             // 分页查询Article列表
	}
}
