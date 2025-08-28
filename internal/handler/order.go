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
