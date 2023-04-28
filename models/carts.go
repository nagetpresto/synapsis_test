package models

type Cart struct {
    ID             int                  `json:"id" gorm:"primary_key:auto_increment"`
    UserID         int                  `json:"user_id"`
    User           UserProfileResponse  `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
    ProductID      int                  `json:"-" gorm:"type: int"`
    Product        Product              `json:"product" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
    Qty            int                  `json:"qty" gorm:"type: int"`
    Amount         int                  `json:"amount" gorm:"type: int"`
    TransactionID  int                  `json:"transaction_id" gorm:"type: int"`
    Transaction    Transaction          `json:"-"`
}