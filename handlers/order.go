package handlers

import (
	"github.com/gin-gonic/gin"
	"go_final/models"
	"go_final/repositories"
	"net/http"
	"strconv"
)

type OrderHandler interface {
	OrderProducts(*gin.Context)
	UpdateOrder(*gin.Context)
	DeleteOrder(*gin.Context)
	DeleteOrderItem(*gin.Context)
	GetCurrentOrder(*gin.Context)
}

type orderHandler struct {
	repo repositories.OrderRepository
}

func NewOrderHandler() OrderHandler {
	return &orderHandler{
		repo: repositories.NewOrderRepository(),
	}
}

func (h *orderHandler) GetCurrentOrder(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	order, err := h.repo.GetCurrentOrder(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (h *orderHandler) OrderProducts(ctx *gin.Context) {
	var input models.CartRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := ctx.GetFloat64("userID")
	if err := h.repo.OrderProducts(uint(userID), input.CartItems); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, input)
}

func (h *orderHandler) UpdateOrder(ctx *gin.Context) {
	var input models.CartRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := ctx.GetFloat64("userID")
	if err := h.repo.UpdateOrder(uint(userID), input.CartItems); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, input)
}

func (h *orderHandler) DeleteOrder(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	if err := h.repo.DeleteOrder(uint(userID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete order"})
		return
	}
	ctx.Status(http.StatusOK)
}

func (h *orderHandler) DeleteOrderItem(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")

	orderItemIDStr := ctx.Param("order_item_id")
	orderItemID, err := strconv.Atoi(orderItemIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order item ID"})
		return
	}

	if err = h.repo.DeleteOrderItem(uint(userID), uint(orderItemID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete order item"})
		return
	}

	ctx.Status(http.StatusOK)
}
