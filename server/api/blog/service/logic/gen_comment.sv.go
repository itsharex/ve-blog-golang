package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type CommentService struct {
	svcCtx *svc.ServiceContext
}

func NewCommentService(svcCtx *svc.ServiceContext) *CommentService {
	return &CommentService{
		svcCtx: svcCtx,
	}
}

// 创建Comment记录
func (s *CommentService) CreateComment(reqCtx *request.Context, comment *entity.Comment) (data *entity.Comment, err error) {
	return s.svcCtx.CommentRepository.CreateComment(reqCtx, comment)
}

// 删除Comment记录
func (s *CommentService) DeleteComment(reqCtx *request.Context, comment *entity.Comment) (rows int64, err error) {
	return s.svcCtx.CommentRepository.DeleteComment(reqCtx, comment)
}

// 更新Comment记录
func (s *CommentService) UpdateComment(reqCtx *request.Context, comment *entity.Comment) (data *entity.Comment, err error) {
	return s.svcCtx.CommentRepository.UpdateComment(reqCtx, comment)
}

// 查询Comment记录
func (s *CommentService) FindComment(reqCtx *request.Context, comment *entity.Comment) (data *entity.Comment, err error) {
	return s.svcCtx.CommentRepository.GetComment(reqCtx, comment.ID)
}

// 批量删除Comment记录
func (s *CommentService) DeleteCommentByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.CommentRepository.DeleteCommentByIds(reqCtx, ids)
}

// 分页获取Comment记录
func (s *CommentService) FindCommentList(reqCtx *request.Context, page *request.PageInfo) (list []*response.CommentDTO, total int64, err error) {
	commentList, total, err := s.svcCtx.CommentRepository.FindCommentList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	var userIds []int
	var commentIds []int
	for _, item := range commentList {
		userIds = append(userIds, item.UserID)
		commentIds = append(commentIds, item.ID)
	}

	// 查询用户
	users, _, _ := s.svcCtx.UserInformationRepository.FindUserInformationList(reqCtx, &request.PageInfo{
		Conditions: []*request.Condition{{
			Field: "id",
			Rule:  "in",
			Value: userIds,
		}},
	})
	var userMap = make(map[int]*entity.UserInformation)
	for _, item := range users {
		userMap[item.ID] = item
	}

	for _, item := range commentList {
		// 查询评论下所有回复列表,只显示五条
		replyList, count, _ := s.FindCommonReplyList(reqCtx, item.ID, &request.PageInfo{
			Page:     1,
			PageSize: 5,
		})
		// 查询当前评论下所有回复列表
		data := &response.CommentDTO{
			ID:             item.ID,
			UserID:         item.UserID,
			CommentContent: item.CommentContent,
			LikeCount:      100,
			CreatedAt:      item.CreatedAt,
			ReplyCount:     count,
			ReplyDTOList:   replyList,
		}

		// 用户信息
		info, _ := userMap[item.UserID]
		if info != nil {
			data.Nickname = info.Nickname
			data.Avatar = info.Avatar
			data.WebSite = info.WebSite
		}

		// 回复的用户信息
		//rinfo, _ := userMap[item.ReplyUserID]
		//if rinfo != nil {
		//	data.ReplyUserID = rinfo.ID
		//	data.ReplyNickname = rinfo.Nickname
		//	data.ReplyWebSite = rinfo.WebSite
		//}

		list = append(list, data)
	}

	return
}

// 查询Comment记录
func (s *CommentService) FindCommonReplyList(reqCtx *request.Context, commentId int, page *request.PageInfo) (list []*response.ReplyDTO, total int64, err error) {
	// 查询评论下所有回复列表
	replyList, total, _ := s.svcCtx.CommentRepository.FindCommentReplyList(commentId, page)

	// 收集需要查询的用户id
	var userIds []int
	for _, item := range replyList {
		userIds = append(userIds, item.UserID)
		userIds = append(userIds, item.ReplyUserID)
	}

	// 查询用户
	users, _, _ := s.svcCtx.UserInformationRepository.FindUserInformationList(reqCtx, &request.PageInfo{
		Conditions: []*request.Condition{{
			Field: "id",
			Rule:  "in",
			Value: userIds,
		}},
	})
	var userMap = make(map[int]*entity.UserInformation)
	for _, item := range users {
		userMap[item.ID] = item
	}

	// 组装返回数据
	for _, item := range replyList {

		data := &response.ReplyDTO{
			ID:             item.ID,
			ParentID:       item.ParentID,
			UserID:         item.UserID,
			ReplyUserID:    item.ReplyUserID,
			CommentContent: item.CommentContent,
			LikeCount:      5,
			CreatedAt:      item.CreatedAt,
		}

		// 用户信息
		info, _ := userMap[item.UserID]
		if info != nil {
			data.Nickname = info.Nickname
			data.Avatar = info.Avatar
			data.WebSite = info.WebSite
		}

		// 回复的用户信息
		rinfo, _ := userMap[item.ReplyUserID]
		if rinfo != nil {
			data.ReplyUserID = rinfo.ID
			data.ReplyNickname = rinfo.Nickname
			data.ReplyWebSite = rinfo.WebSite
		}

		list = append(list, data)
	}
	return list, total, nil
}

// 查询Comment后台记录
func (s *CommentService) FindCommonBackList(reqCtx *request.Context, page *request.PageInfo) (list []*response.CommentBackDTO, total int64, err error) {
	// 使用用户昵称查询
	username := page.FindCondition("username")
	if username != nil {
		accounts, _, err := s.svcCtx.UserAccountRepository.FindUserAccountList(reqCtx, &request.PageInfo{
			Page:       0,
			PageSize:   0,
			Orders:     nil,
			Conditions: []*request.Condition{username},
		})
		if err != nil {
			return nil, 0, err
		}

		var userIds []int
		for _, item := range accounts {
			userIds = append(userIds, item.ID)
		}
		// 替换查询条件
		username.Field = "user_id"
		username.Value = userIds
		username.Rule = "in"
	}

	// 查询评论下所有回复列表
	replyList, total, _ := s.svcCtx.CommentRepository.FindCommentList(reqCtx, page)

	// 收集需要查询的用户id
	var userIds []int
	for _, item := range replyList {
		userIds = append(userIds, item.UserID)
		userIds = append(userIds, item.ReplyUserID)
	}

	// 查询用户
	users, _, _ := s.svcCtx.UserInformationRepository.FindUserInformationList(reqCtx, &request.PageInfo{
		Conditions: []*request.Condition{{
			Field: "id",
			Rule:  "in",
			Value: userIds,
		}},
	})
	var userMap = make(map[int]*entity.UserInformation)
	for _, item := range users {
		userMap[item.ID] = item
	}

	// 组装返回数据
	for _, item := range replyList {

		data := &response.CommentBackDTO{
			ID:             0,
			Avatar:         "",
			Nickname:       "",
			ReplyNickname:  "",
			ArticleTitle:   "测试",
			CommentContent: item.CommentContent,
			Type:           item.Type,
			IsReview:       item.IsReview,
			CreatedAt:      item.CreatedAt,
		}

		// 用户信息
		info, _ := userMap[item.UserID]
		if info != nil {
			data.Nickname = info.Nickname
			data.Avatar = info.Avatar
		}

		// 回复的用户信息
		rinfo, _ := userMap[item.ReplyUserID]
		if rinfo != nil {
			data.ReplyNickname = rinfo.Nickname
		}

		list = append(list, data)
	}
	return list, total, nil
}

// 点赞Comment
func (s *CommentService) LikeComment(reqCtx *request.Context, commentId int) (data interface{}, err error) {

	return s.svcCtx.CommentRepository.LikeComment(reqCtx, reqCtx.UID, commentId)
}