package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/controller/svc"
	entity2 "github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	request2 "github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type UserController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewUserController(svcCtx *svc.ControllerContext) *UserController {
	return &UserController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		User
// @Summary		获取用户菜单权限
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/user/menus [get]
func (s *UserController) GetUserMenus(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.MenuService.GetUserMenus(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		User
// @Summary		获取用户接口权限
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/user/apis [get]
func (s *UserController) GetUserApis(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.ApiService.GetUserApis(reqCtx, nil)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		User
// @Summary		获取用户登录历史
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageQuery	true	"分页参数"
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/user/login_history [get]
func (s *UserController) GetLoginHistory(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request2.PageQuery
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.UserService.GetLoginHistory(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.Limit(),
	})
}

// @Tags		User
// @Summary		获取用户信息
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/user/info [get]
func (s *UserController) GetUserInfo(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.GetUserInfo(reqCtx, reqCtx.UID)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		User
// @Summary		修改用户信息
// @Security	ApiKeyAuth
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		data	body		entity.UserInformation					true	"请求body"
// @Success		200		{object}	response.Response{data=entity.Upload}	"返回信息"
// @Router		/user/info [post]
func (s *UserController) UpdateUserInfo(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req entity2.UserInformation
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.UpdateUserInfo(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		User
// @Summary		更换用户头像
// @Security	ApiKeyAuth
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		data	body		entity.Upload							true	"请求body"
// @Success		200		{object}	response.Response{data=entity.Upload}	"返回信息"
// @Router		/user/avatar [post]
func (s *UserController) UpdateUserAvatar(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.UpdateUserAvatar(reqCtx, file)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		User
// @Summary		获取用户列表
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageQuery	true	"分页参数"
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/admin/users [post]
func (s *UserController) GetUserList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request2.PageQuery
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.UserService.GetUserList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.Limit(),
	})
}

// @Tags		User
// @Summary		获取用户地区
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Param		page	body		request.PageQuery	true	"分页参数"
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/admin/user/areas [post]
func (s *UserController) GetUserAreas(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request2.PageQuery
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.UserService.GetUserAreas(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.Limit(),
	})
}

// @Tags		User
// @Summary		修改用户状态
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Param		data	body		entity.UserAccount	true	"请求数据"
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/admin/user/update_status [post]
func (s *UserController) UpdateUserStatus(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req entity2.UserAccount
	err = s.ShouldBindJSON(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.UpdateUserStatus(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		User
// @Summary		修改用户角色
// @Security	ApiKeyAuth
// @Accept		application/json
// @Produce		application/json
// @Param		data	body		request.UpdateUserRoles				true	"请求数据"
// @Success		200		{object}	response.Response{data=entity.Role}	"返回信息"
// @Router		/admin/user/update_roles [post]
func (s *UserController) UpdateUserRoles(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request2.UpdateUserRoles
	err = s.ShouldBindJSON(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.UserService.UpdateUserRoles(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}