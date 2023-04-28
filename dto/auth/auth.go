package authdto

type AuthRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Status 	 string `json:"status" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ProfileUpdateRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Image	 string `json:"image" form:"image"`
}

type LoginResponse struct {
	Name     	string 	`gorm:"type: varchar(255)" json:"name"`
	Email    	string 	`gorm:"type: varchar(255)" json:"email"`
	Password 	string 	`gorm:"type: varchar(255)" json:"password"`
	Image 		string 	`gorm:"type: varchar(255)" json:"image"`
	Status 		string 	`gorm:"type: varchar(255)" json:"status"`
	IsConfirmed	bool 	`gorm:"type:boolean" json:"is_confirmed"`
	ConfirmCode	string 	`gorm:"type: varchar(255)" json:"confirm_code"`
	Token    	string 	`gorm:"type: varchar(255)" json:"token"`
}

type ProfileResponse struct {
	ID       	int    	`json:"id" gorm:"primary_key:auto_increment"`
	Name     	string 	`gorm:"type: varchar(255)" json:"name"`
	Email    	string 	`gorm:"type: varchar(255)" json:"email"`
	Password 	string 	`gorm:"type: varchar(255)" json:"password"`
	Image 		string 	`gorm:"type: varchar(255)" json:"image"`
	Status 		string 	`gorm:"type: varchar(255)" json:"status"`
	IsConfirmed	bool 	`gorm:"type:boolean" json:"is_confirmed"`
	ConfirmCode	string 	`gorm:"type: varchar(255)" json:"confirm_code"`
}