package router

import (
	"worker-api/internal/handler"
	"worker-api/internal/service"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	orderService := service.NewOrderService(db)
	orderHandler := handler.NewOrderHandler(orderService)

	orders := e.Group("/orders")
	orders.POST("/", orderHandler.CreateOrder).Name = "create-order"
	orders.GET("/:id", orderHandler.GetOrderStatus).Name = "get-order-status"
}
