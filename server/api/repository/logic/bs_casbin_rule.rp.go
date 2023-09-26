package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type CasbinRuleRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewCasbinRuleRepository(svcCtx *svc.RepositoryContext) *CasbinRuleRepository {
	return &CasbinRuleRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建CasbinRule记录
func (s *CasbinRuleRepository) CreateCasbinRule(ctx context.Context, casbinRule *entity.CasbinRule) (out *entity.CasbinRule, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Create(&casbinRule).Error
	if err != nil {
		return nil, err
	}
	return casbinRule, err
}

// 删除CasbinRule记录
func (s *CasbinRuleRepository) DeleteCasbinRule(ctx context.Context, casbinRule *entity.CasbinRule) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx)
	query := db.Delete(&casbinRule)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新CasbinRule记录
func (s *CasbinRuleRepository) UpdateCasbinRule(ctx context.Context, casbinRule *entity.CasbinRule) (out *entity.CasbinRule, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Save(&casbinRule).Error
	if err != nil {
		return nil, err
	}
	return casbinRule, err
}

// 查询CasbinRule记录
func (s *CasbinRuleRepository) FindCasbinRule(ctx context.Context, id int) (out *entity.CasbinRule, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除CasbinRule记录
func (s *CasbinRuleRepository) DeleteCasbinRuleByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx)
	query := db.Delete(&[]entity.CasbinRule{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询CasbinRule记录
func (s *CasbinRuleRepository) FindCasbinRuleList(ctx context.Context, page *request.PageQuery) (list []*entity.CasbinRule, total int64, err error) {
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
