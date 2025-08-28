package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"worker-api/config"
	"worker-api/internal/domain"
	"worker-api/pkg/email"

	"gorm.io/gorm"
)

type OrderWorker struct {
	db    *gorm.DB
	email *email.Service
}

func NewOrderWorker(db *gorm.DB, email *email.Service) *OrderWorker {
	return &OrderWorker{
		db:    db,
		email: email,
	}
}

func (w *OrderWorker) Start(ctx context.Context) {
	log.Println("Starting order worker...")

	for {
		select {
		case <-ctx.Done():
			return
		default:
			w.processMessages(ctx)
		}
	}
}

func (w *OrderWorker) processMessages(ctx context.Context) {
	result, err := config.ReceiveMessageFromSQS(ctx)
	if err != nil {
		log.Printf("Error receiving message: %v", err)
		return
	}

	for _, message := range result.Messages {
		var order domain.Order
		if err := json.Unmarshal([]byte(*message.Body), &order); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		if err := w.validateOrder(&order); err != nil {
			log.Printf("Error validating order %d: %v", order.ID, err)
			continue
		}

		order.Status = "paid"
		if err := w.db.Save(&order).Error; err != nil {
			log.Printf("Error updating order %d: %v", order.ID, err)
			continue
		}

		if err := w.email.SendOrderConfirmation(order.CustomerEmail, order.ID); err != nil {
			log.Printf("Error sending email for order %d: %v", order.ID, err)
		}

		if err := config.DeleteMessageFromSQS(ctx, message.ReceiptHandle); err != nil {
			log.Printf("Error deleting message: %v", err)
		}
	}
}

func (w *OrderWorker) validateOrder(order *domain.Order) error {
	if order.Amount <= 0 {
		return fmt.Errorf("invalid amount")
	}
	if order.CustomerID == 0 {
		return fmt.Errorf("invalid customer")
	}
	return nil
}
