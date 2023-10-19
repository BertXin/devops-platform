package repository

import (
	"context"
	"devops-platform/internal/deploy-system/server/internal/domain"
	"devops-platform/pkg/types"
)

/**
 * 根据ID获取镜像仓库信息
 */
func (r *Repository) GetImageRepositoryByID(ctx context.Context, ID types.Long) (imageRepository *domain.ImageRepository, err error) {
	imageRepository = new(domain.ImageRepository)

	err = r.DB(ctx).Take(imageRepository, ID).Error
	if imageRepository.ID == 0 {
		imageRepository = nil
	}
	return
}

/**
 * 创建镜像仓库信息
 */
func (r *Repository) CreateImageRepository(ctx context.Context, imageRepository *domain.ImageRepository) (err error) {
	result := r.DB(ctx).Create(imageRepository)
	err = result.Error
	return
}

/**
 * 保存、更新仓库信息
 */
func (r *Repository) SaveImageRepository(ctx context.Context, imageRepository *domain.ImageRepository) (err error) {
	result := r.DB(ctx).Save(imageRepository)
	err = result.Error
	return
}

/**
 * 根据镜像仓库名称查找镜像仓库信息
 */
func (r *Repository) FindImageRepositoryByNameAndAddress(ctx context.Context, name string, address string, pagination types.Pagination) (repository []domain.ImageRepository, total int64, err error) {
	err = r.DB(ctx).Model(&domain.ImageRepository{}).Where("name like ?", "%"+name+"%").Where("address like ?", "%"+address+"%").Count(&total).Error
	if total == 0 || err != nil {
		return
	}
	err = r.DB(ctx).Limit(pagination.Limit()).Offset(pagination.Offset()).Where("name like ?", "%"+name+"%").Where("address like ?", "%"+address+"%").Find(&repository).Error
	return
}

/*
 * 删除构建服务器
 */
func (r *Repository) DeleteImageRepositoryById(ctx context.Context, id types.Long) (err error) {
	imageRepository := new(domain.ImageRepository)
	imageRepository.ID = id
	err = r.DB(ctx).Delete(imageRepository).Error
	return err
}
