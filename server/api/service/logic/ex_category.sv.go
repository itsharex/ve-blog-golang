package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// 分页获取Category记录
func (s *CategoryService) FindCategoryDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.CategoryDetailsDTO, total int64, err error) {
	categories, total, err := s.svcCtx.CategoryRepository.FindCategoryList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	// 查询分类下的文章数量

	for _, in := range categories {

		_, articleCount, err := s.svcCtx.ArticleRepository.FindArticleListByCategoryId(reqCtx, in.ID)
		if err != nil {
			return nil, 0, err
		}

		out := &response.CategoryDetailsDTO{
			ID:           in.ID,
			CategoryName: in.CategoryName,
			ArticleCount: articleCount,
			CreatedAt:    in.CreatedAt,
			UpdatedAt:    in.UpdatedAt,
		}

		list = append(list, out)
	}

	return list, total, err
}
