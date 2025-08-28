package domain

type Order struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	CustomerID    uint    `json:"customer_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	CustomerEmail string  `json:"customer_email"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type OrderRequest struct {
	CustomerID    uint    `json:"customer_id" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
	CustomerEmail string  `json:"customer_email" validate:"required,email"`
}
