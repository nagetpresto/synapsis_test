package models

type Cart struct {
    ID             int                  `json:"id" gorm:"primary_key:auto_increment"`
    UserID         int                  `json:"user_id"`
    User           UserProfileResponse  `json:"-"`
    ProductID      int                  `json:"product_id" gorm:"type: int"`
    Product        Product              `json:"product"`
    Qty            int                  `json:"qty" gorm:"type: int"`
    Amount         int                  `json:"amount" gorm:"type: int"`
    TransactionID  int                  `json:"transaction_id" gorm:"type: int"`
    Transaction    Transaction          `json:"-"`
}