package logic

import (
	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/repository/svc"
	"gorm.io/gorm"
)

type PhotoRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewPhotoRepository(svcCtx *svc.RepositoryContext) *PhotoRepository {
	return &PhotoRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Photo记录
func (s *PhotoRepository) CreatePhoto(photo *entity.Photo) (out *entity.Photo, err error) {
	db := s.DbEngin
	err = db.Create(&photo).Error
	if err != nil {
		return nil, err
	}
	return photo, err
}

// 删除Photo记录
func (s *PhotoRepository) DeletePhoto(photo *entity.Photo) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&photo)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Photo记录
func (s *PhotoRepository) UpdatePhoto(photo *entity.Photo) (out *entity.Photo, err error) {
	db := s.DbEngin
	err = db.Save(&photo).Error
	if err != nil {
		return nil, err
	}
	return photo, err
}

// 查询Photo记录
func (s *PhotoRepository) GetPhoto(id int) (out *entity.Photo, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Photo记录
func (s *PhotoRepository) DeletePhotoByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Photo{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询Photo记录
func (s *PhotoRepository) FindPhotoList(page *request.PageInfo) (list []*entity.Photo, total int64, err error) {
	// 创建db
	db := s.DbEngin

	// 如果有搜索条件
	if len(page.Conditions) != 0 {
		query, args := page.WhereClause()
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(page.Orders) != 0 {
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
