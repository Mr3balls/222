package usecase

import (
	"context"
	"order-service/internal/models"
	"order-service/internal/repository"
)

type OrderUseCase struct {
	orderRepo *repository.OrderRepository
}

func NewOrderUseCase(orderRepo *repository.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
	}
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, order *models.Order) error {
	return uc.orderRepo.CreateOrder(ctx, order)
}

func (uc *OrderUseCase) GetOrderByID(ctx context.Context, id string) (*models.Order, error) {
	return uc.orderRepo.GetOrderByID(ctx, id)
}

func (uc *OrderUseCase) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	return uc.orderRepo.GetAllOrders(ctx)
}

func (uc *OrderUseCase) UpdateOrder(ctx context.Context, id string, order *models.Order) error {
	return uc.orderRepo.UpdateOrder(ctx, id, order)
}
