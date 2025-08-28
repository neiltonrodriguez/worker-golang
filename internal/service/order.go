package service

import (
	"context"
	"encoding/json"
	"worker-api/config"
	"worker-api/internal/domain"

	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) CreateOrder(ctx context.Context, request *domain.OrderRequest) error {
	// Create order with waiting status
	order := &domain.Order{
		CustomerID:    request.CustomerID,
		Amount:        request.Amount,
		CustomerEmail: request.CustomerEmail,
		Status:        "waiting",
	}

	// Save to database
	if err := s.db.Create(order).Error; err != nil {
		return err
	}

	// Send to SQS
	messageBody, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return config.SendMessageToSQS(ctx, string(messageBody))
}
