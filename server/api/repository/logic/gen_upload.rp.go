package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type UploadRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewUploadRepository(svcCtx *svc.RepositoryContext) *UploadRepository {
	return &UploadRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Upload记录
func (s *UploadRepository) CreateUpload(ctx context.Context, upload *entity.Upload) (out *entity.Upload, err error) {
	db := s.DbEngin
	err = db.Create(&upload).Error
	if err != nil {
		return nil, err
	}
	return upload, err
}

// 删除Upload记录
func (s *UploadRepository) DeleteUpload(ctx context.Context, upload *entity.Upload) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&upload)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Upload记录
func (s *UploadRepository) UpdateUpload(ctx context.Context, upload *entity.Upload) (out *entity.Upload, err error) {
	db := s.DbEngin
	err = db.Save(&upload).Error
	if err != nil {
		return nil, err
	}
	return upload, err
}

// 查询Upload记录
func (s *UploadRepository) FindUpload(ctx context.Context, id int) (out *entity.Upload, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Upload记录
func (s *UploadRepository) DeleteUploadByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Upload{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询Upload记录
func (s *UploadRepository) FindUploadList(ctx context.Context, page *request.PageQuery) (list []*entity.Upload, total int64, err error) {
	// 创建db
	db := s.DbEngin

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