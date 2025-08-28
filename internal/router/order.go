package router

import (
	"worker-api/internal/handler"
	"worker-api/internal/service"
	"worker-api/pkg/database"

	"github.com/labstack/echo"
)

func RegisterRoutes(e *echo.Echo, db *database.DB) {
	// Order routes
	orderService := service.NewOrderService(db)
	orderHandler := handler.NewOrderHandler(orderService)

	orders := e.Group("/orders")
	orders.POST("/", orderHandler.CreateOrder).Name = "create-order"
	orders.GET("/:id", orderHandler.GetOrderStatus).Name = "get-order-status"
}
