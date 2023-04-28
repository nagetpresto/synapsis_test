package repositories

import (
	"BE/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllTransaction() ([]models.Transaction, error)
	GetOneTransaction(ID int) (models.Transaction, error)
	GetUserTrans(UserID int) ([]models.Transaction, error)

	GetActiveTransaction(UserID int) (models.Transaction, error)
	DoTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	UpdateStatsTransaction(status string, ID int) error //for midtrans

	UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	GetProduct() (models.Product, error)
	UpdateProductStock(product models.Product) (models.Product, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GetAllTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Order("id").Preload("User").Preload("Cart").Preload("Cart.Product").Preload("Cart.Product.Category").Find(&transactions, "status!=?", "active").Error
	return transactions, err
}

func (r *repository) GetOneTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Order("id").Preload("User").Preload("Cart").Preload("Cart.Product").Preload("Cart.Product.Category").First(&transaction, ID).Error
	return transaction, err
}

func (r *repository) GetUserTrans(UserID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Order("id").Preload("User").Preload("Cart").Preload("Cart.Product").Preload("Cart.Product.Category").Where("user_id = ? AND status!=?", UserID, "active").Order("id desc").Find(&transaction).Error
	return transaction, err
}

func (r *repository) GetActiveTransaction(UserID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Order("id").Preload("User").Preload("Cart").Preload("Cart.Product").Preload("Cart.Product.Category").Where("user_id = ? AND status = ?", UserID, "active").First(&transaction).Error
	return transaction, err
}

func (r *repository) DoTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error
	return transaction, err
}

func (r *repository) UpdateStatsTransaction(status string, ID int) error {
	var transaction models.Transaction
	r.db.Preload("User").Preload("Cart").Preload("Cart.Product").Preload("Cart.Product.Category").First(&transaction, ID)

	// if status == "Success" {
	// 	var product models.Product
	// 	r.db.First(&product, transaction.ID)
	// 	for _, cart := range transaction.Cart {
    //         product := cart.Product
    //         product.Stock -= cart.Qty
    //         if err := r.db.Save(&product).Error; err != nil {
    //             return err
    //         }
    //     }
	// }

	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return err
}

func (r *repository) UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error
	return transaction, err
}

func (r *repository) GetProduct() (models.Product, error) {
	var product models.Product
	err := r.db.First(&product).Error
	return product, err
}

func (r *repository) UpdateProductStock(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}