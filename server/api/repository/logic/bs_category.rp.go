package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type CategoryRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewCategoryRepository(svcCtx *svc.RepositoryContext) *CategoryRepository {
	return &CategoryRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Category记录
func (s *CategoryRepository) CreateCategory(ctx context.Context, category *entity.Category) (out *entity.Category, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Create(&category).Error
	if err != nil {
		return nil, err
	}
	return category, err
}

// 更新Category记录
func (s *CategoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) (out *entity.Category, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Save(&category).Error
	if err != nil {
		return nil, err
	}
	return category, err
}

// 删除Category记录
func (s *CategoryRepository) DeleteCategory(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx)
	query := db.Delete(&entity.Category{}, "id = ?", id)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 查询Category记录
func (s *CategoryRepository) FindCategory(ctx context.Context, id int) (out *entity.Category, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Category记录
func (s *CategoryRepository) DeleteCategoryByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx)
	query := db.Delete(&entity.Category{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询Category记录
func (s *CategoryRepository) FindCategoryList(ctx context.Context, page *request.PageQuery) (list []*entity.Category, total int64, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	// 如果有搜索条件
	if len(page.Conditions) != 0 {
		query, args := page.WhereClause()
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(page.Sorts) != 0 {
		db = db.Order(page.OrderClause())
	}

	// 查询总数,要在使用limit之前
	err = db.Model(&list).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 如果有分页参数
	if page.Page != 0 || page.PageSize != 0 {
		limit := page.Limit()
		offset := page.Offset()
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
