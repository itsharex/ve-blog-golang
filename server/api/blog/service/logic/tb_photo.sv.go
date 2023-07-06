package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type PhotoService struct {
	svcCtx *svc.ServiceContext
}

func NewPhotoService(svcCtx *svc.ServiceContext) *PhotoService {
	return &PhotoService{
		svcCtx: svcCtx,
	}
}

// 创建Photo记录
func (s *PhotoService) CreatePhoto(reqCtx *request.Context, photo *entity.Photo) (data *entity.Photo, err error) {
	return s.svcCtx.PhotoRepository.CreatePhoto(photo)
}

// 删除Photo记录
func (s *PhotoService) DeletePhoto(reqCtx *request.Context, photo *entity.Photo) (rows int64, err error) {
	return s.svcCtx.PhotoRepository.DeletePhoto(photo)
}

// 更新Photo记录
func (s *PhotoService) UpdatePhoto(reqCtx *request.Context, photo *entity.Photo) (data *entity.Photo, err error) {
	return s.svcCtx.PhotoRepository.UpdatePhoto(photo)
}

// 查询Photo记录
func (s *PhotoService) GetPhoto(reqCtx *request.Context, photo *entity.Photo) (data *entity.Photo, err error) {
	return s.svcCtx.PhotoRepository.GetPhoto(photo.ID)
}

// 批量删除Photo记录
func (s *PhotoService) DeletePhotoByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.PhotoRepository.DeletePhotoByIds(ids)
}

// 分页获取Photo记录
func (s *PhotoService) FindPhotoList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.Photo, total int64, err error) {
	return s.svcCtx.PhotoRepository.FindPhotoList(page)
}
