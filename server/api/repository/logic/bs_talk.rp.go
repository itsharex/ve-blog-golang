package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type TalkRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewTalkRepository(svcCtx *svc.RepositoryContext) *TalkRepository {
	return &TalkRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Talk记录
func (s *TalkRepository) CreateTalk(ctx context.Context, talk *entity.Talk, conditions ...*request.Condition) (out *entity.Talk, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Create(&talk).Error
	if err != nil {
		return nil, err
	}
	return talk, err
}

// 更新Talk记录
func (s *TalkRepository) UpdateTalk(ctx context.Context, talk *entity.Talk, conditions ...*request.Condition) (out *entity.Talk, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Save(&talk).Error
	if err != nil {
		return nil, err
	}
	return talk, err
}

// 删除Talk记录
func (s *TalkRepository) DeleteTalk(ctx context.Context, id int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.Talk{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询Talk记录
func (s *TalkRepository) FindTalk(ctx context.Context, id int, conditions ...*request.Condition) (out *entity.Talk, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Talk记录
func (s *TalkRepository) DeleteTalkByIds(ctx context.Context, ids []int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.Talk{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 分页查询Talk记录
func (s *TalkRepository) FindTalkList(ctx context.Context, page *request.PageQuery) (list []*entity.Talk, err error) {
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

	// 如果有分页参数
	if page.Page != 0 || page.PageSize != 0 {
		limit := page.Limit()
		offset := page.Offset()
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 查询总数
func (s *TalkRepository) Count(ctx context.Context, conditions ...*request.Condition) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Model(&entity.Talk{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
