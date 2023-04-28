package models

type Product struct {
	ID          int    		`json:"id" gorm:"primary_key:auto_increment"`
	CategoryID  int    		`json:"-"`
    Category    Category	`json:"category" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	Name        string 		`json:"name" gorm:"type: varchar(255)"`
	Stock       int    		`json:"stock" gorm:"type: int"`
	Price       int    		`json:"price" gorm:"type: int"`
	Description string 		`json:"description" gorm:"type: text"`
	Image       string 		`json:"image" gorm:"type: varchar(255)"`
}