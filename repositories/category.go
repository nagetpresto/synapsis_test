package repositories

import (
	"BE/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategory() ([]models.Category, error)
	GetOneCategory(ID int) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	DeleteCategory(category models.Category, ID int) (models.Category, error)
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllCategory() ([]models.Category, error) {
	var category []models.Category
	err := r.db.Order("id").Find(&category).Error

	return category, err
}

func (r *repository) GetOneCategory(ID int) (models.Category, error) {
	var category models.Category
	err := r.db.First(&category, ID).Error

	return category, err
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error

	return category, err
}

func (r *repository) UpdateCategory(category models.Category) (models.Category, error) {
	err := r.db.Save(&category).Error

	return category, err
}

func (r *repository) DeleteCategory(category models.Category, ID int) (models.Category, error) {
	err := r.db.Delete(&category, ID).Scan(&category).Error

	return category, err
}