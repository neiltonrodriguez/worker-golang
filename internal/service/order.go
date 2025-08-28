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
	order := &domain.Order{
		CustomerID:    request.CustomerID,
		Amount:        request.Amount,
		CustomerEmail: request.CustomerEmail,
		Status:        "waiting",
	}

	if err := s.db.Create(order).Error; err != nil {
		return err
	}

	messageBody, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return config.SendMessageToSQS(ctx, string(messageBody))
}

func (s *OrderService) GetOrder(ctx context.Context, orderID string) (*domain.Order, error) {
	var order domain.Order
	result := s.db.First(&order, orderID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &order, nil
}
