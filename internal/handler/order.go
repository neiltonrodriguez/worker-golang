package handler

import (
	"net/http"
	"worker-api/internal/domain"
	"worker-api/internal/service"

	"github.com/labstack/echo"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var request domain.OrderRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if err := c.Validate(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err := h.orderService.CreateOrder(c.Request().Context(), &request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to process order",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Pedido em análise",
	})
}

func (h *OrderHandler) GetOrderStatus(c echo.Context) error {
	orderId := c.Param("id")
	if orderId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Order ID is required",
		})
	}

	order, err := h.orderService.GetOrder(c.Request().Context(), orderId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get order status",
		})
	}

	if order == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Order not found",
		})
	}

	return c.JSON(http.StatusOK, order)
}
