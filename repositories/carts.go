package repositories

import (
	"BE/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetAllCart() ([]models.Cart, error)
	GetOneCart(ID int) (models.Cart, error)
	GetActiveCart(TransID int) ([]models.Cart, error)

	GetUserEmailStats(ID int) (models.User, error)
	CreateCart(Cart models.Cart) (models.Cart, error)
	GetActiveTrans(UserID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetProd(ID int) (models.Product, error)
	GetActiveProduct(UserID int, TransID int, ProductID int) (models.Cart, error)

	UpdateCart(Cart models.Cart, ID int) (models.Cart, error)
	DeleteCart(Cart models.Cart, ID int) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllCart() ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Order("id").Preload("Product").Preload("Product.Category").Find(&cart).Error

	return cart, err
}

func (r *repository) GetOneCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Product").Preload("Product.Category").First(&cart, ID).Error

	return cart, err
}

func (r *repository) GetActiveCart(TransID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Order("id").Preload("Product").Preload("Product.Category").Find(&carts, "transaction_id = ?", TransID).Error

	return carts, err
}

func (r *repository) GetUserEmailStats(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Transaction").Preload("Transaction.User").Preload("Transaction.Cart").Preload("Transaction.Cart.Product").First(&user, "id=? AND is_confirmed=?", ID, true).Error // add this code
	return user, err
}

func (r *repository) GetActiveTrans(UserID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Where("user_id = ? AND status = ?", UserID, "active").First(&transaction).Error
	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *repository) GetProd(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ?", ID).First(&product).Error // add this code

	return product, err
}

func (r *repository) GetActiveProduct(UserID int, TransID int,ProductID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("user_id = ? AND transaction_id = ? AND product_id = ?", UserID, TransID, ProductID).First(&cart).Error
	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error
	return cart, err
}

func (r *repository) UpdateCart(Cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Save(&Cart).Error

	return Cart, err
}

func (r *repository) DeleteCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}