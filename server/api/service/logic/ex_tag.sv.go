package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// 分页获取Tag记录
func (s *TagService) FindTagDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.TagDetailsDTO, total int64, err error) {
	categories, total, err := s.svcCtx.TagRepository.FindTagList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	// 查询分类下的文章数量

	for _, in := range categories {

		_, articleCount, err := s.svcCtx.ArticleRepository.FindArticleListByTagId(reqCtx, in.ID)
		if err != nil {
			return nil, 0, err
		}

		out := &response.TagDetailsDTO{
			ID:           in.ID,
			TagName:      in.TagName,
			ArticleCount: articleCount,
			CreatedAt:    in.CreatedAt,
			UpdatedAt:    in.UpdatedAt,
		}

		list = append(list, out)
	}

	return list, total, err
}