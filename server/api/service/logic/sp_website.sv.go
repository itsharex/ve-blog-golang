package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type WebsiteService struct {
	svcCtx *svc.ServiceContext
}

func NewWebsiteService(svcCtx *svc.ServiceContext) *WebsiteService {
	return &WebsiteService{
		svcCtx: svcCtx,
	}
}

func (s *WebsiteService) GetWebsiteAdminHomeInfo(reqCtx *request.Context, data interface{}) (resp *response.WebsiteAdminHomeInfo, err error) {
	page := &request.PageQuery{}
	// 查询消息数量
	msgCount, err := s.svcCtx.RemarkRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, err
	}

	// 查询用户数量
	userCount, err := s.svcCtx.UserAccountRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, err
	}

	// 查询文章数量
	articles, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
	if err != nil {
		return nil, err
	}

	// 查询分类数量
	categories, err := s.svcCtx.CategoryRepository.FindCategoryList(reqCtx, page)
	if err != nil {
		return nil, err
	}

	// 查询标签数量
	tags, err := s.svcCtx.TagRepository.FindTagList(reqCtx, page)
	if err != nil {
		return nil, err
	}

	uniqueViews, err := s.svcCtx.UniqueViewRepository.FindUniqueViewList(reqCtx, page)
	if err != nil {
		return nil, err
	}

	articleCount, err := s.svcCtx.ArticleRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, err
	}
	resp = &response.WebsiteAdminHomeInfo{
		ViewsCount:            10,
		MessageCount:          msgCount,
		UserCount:             userCount,
		ArticleCount:          articleCount,
		CategoryDTOList:       convertCategoryList(categories),
		TagDTOList:            convertTagList(tags),
		ArticleStatisticsList: convertArticleStatisticsList(articles),
		UniqueViewDTOList:     convertUniqueViewList(uniqueViews),
		ArticleRankDTOList:    convertArticleRankList(articles),
	}

	return resp, err
}
