package models

type Category struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Name        string `json:"name" gorm:"type: varchar(255)"`
	Image       string `json:"image" gorm:"type: varchar(255)"`
}