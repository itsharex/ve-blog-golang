package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/ws"

	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
)

type BlogController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewBlogController(svcCtx *svc.ControllerContext) *BlogController {
	return &BlogController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Blog
// @Summary		创建api路由
// @Description	描述,可以有多个。https://www.jianshu.com/p/4bb4283632e4
// @Security	ApiKeyUser
// @Param		file	formData	file								true	"上传文件"
// @Param		id		path		int									true	"id"
// @Param		token	header		string								true	"token"
// @Param		data	body		entity.Api							true	"创建api路由"
// @Success		200		{object}	response.Response{data=entity.Api}	"返回信息"
// @Router		/version [get]
func (s *BlogController) ApiVersion(c *gin.Context) {

	s.ResponseOk(c, nil)
}

// @Tags		Blog
// @Summary		查询聊天记录
// @Router		/ws [get]
func (s *BlogController) WebSocket(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	// 接收消息
	receive := func(msg []byte) {
		global.LOG.Println(string(msg))

		var chat entity.ChatRecord
		err = jsonconv.JsonToObject(string(msg), &chat)
		if err != nil {
			global.LOG.Error(err)
		}

		if chat.Content == "" {
			return
		}
		if reqCtx.UID != 0 {
			chat.UserID = reqCtx.UID
		}

		_, err = s.svcCtx.ChatRecordService.CreateChatRecord(reqCtx, &chat)
		if err != nil {
			global.LOG.Error(err)
		}
	}

	ws.HandleWebSocket(c.Writer, c.Request, receive)
}

// @Tags		Admin
// @Summary		获取用户地区
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/admin/home [post]
func (s *AdminController) GetHomeInfo(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.BlogService.GetAdminHomeInfo(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Blog
// @Summary		关于我
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Success		200	{object}	response.Response{data=entity.Api}	"返回信息"
// @Router		/about [get]
func (s *BlogController) GetAboutMe(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.WebsiteConfigService.GetAboutMe(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Admin
// @Summary		更新我的信息
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/admin/about [post]
func (s *AdminController) UpdateAboutMe(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req string
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.WebsiteConfigService.UpdateAboutMe(reqCtx, req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Blog
// @Summary		查询聊天记录
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.PageInfo							true	"分页信息"
// @Success		200		{object}	response.Response{data=entity.ChatRecord}	"返回信息"
// @Router		/chat/records [post]
func (s *BlogController) FindChatRecords(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.ChatRecordService.FindChatRecordList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     1,
		PageSize: int(total),
	})
}
