package transactionsdto

type CreateTransactionRequest struct {
	ID			int		`json:"id"`
	UserID   	int    	`json:"user_id"`
	Name	 	string 	`json:"name"`
	Address		string 	`json:"address"`
	PostalCode	string	`json:"postal_code"`
	Phone    	string 	`json:"phone"`
	Status 	 	string 	`json:"status"`
	TotalAmount int 	`json:"total_amount"`
}

type UpdateTransactionRequest struct {
	Status 	 	string 	`json:"status"`
}

type TransactionResponse struct {
	ID     		int    	`json:"id"`
	UserID   	int    	`json:"user_id"`
	Name	 	string 	`json:"name"`
	Address		string 	`json:"address"`
	PostalCode	string	`json:"postal_code"`
	Phone    	string 	`json:"phone"`
	Day  		string 	`json:"day"`
	Date  		string 	`json:"date"`
	Status 	 	string 	`json:"status"`
	TotalAmount int 	`json:"total_amount"`
}