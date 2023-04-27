package repositories

import (
	"BE/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct() ([]models.Product, error)
	GetOneProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(user models.Product) (models.Product, error)
	DeleteProduct(user models.Product, ID int) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllProduct() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Order("id").Find(&products).Error

	return products, err
}

func (r *repository) GetOneProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ID).Error

	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error

	return product, err
}

func (r *repository) DeleteProduct(product models.Product, ID int) (models.Product, error) {
	err := r.db.Delete(&product, ID).Scan(&product).Error

	return product, err
}