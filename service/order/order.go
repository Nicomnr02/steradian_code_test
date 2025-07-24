package orderservice

import (
	"context"
	"steradian_code_test/domain"
	"steradian_code_test/exception"
	carrepository "steradian_code_test/repository/car"
	orderrepository "steradian_code_test/repository/order"
)

type Service interface {
	Create(ctx context.Context, request domain.Order) error
	Update(ctx context.Context, request domain.Order) error
	Delete(ctx context.Context, ID int) error
	GetAll(ctx context.Context) ([]domain.Order, error)
	GetByIDs(ctx context.Context, IDs []int) ([]domain.Order, error)
}

type ServiceImpl struct {
	orderRepo orderrepository.Repository
	carRepo   carrepository.Repository
}

func New(orderRepo orderrepository.Repository, carRepo carrepository.Repository) Service {
	return &ServiceImpl{
		orderRepo: orderRepo,
		carRepo:   carRepo,
	}
}

func (s *ServiceImpl) Create(ctx context.Context, request domain.Order) error {
	err := s.ValidateOrder(ctx, request)
	if err != nil {
		return err
	}

	err = s.orderRepo.Create(ctx, request)
	if err != nil {
		return exception.ErrInternalServer("Failed to create order")
	}

	return nil
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.Order) error {

	err := s.ValidateOrder(ctx, request)
	if err != nil {
		return err
	}

	err = s.orderRepo.Update(ctx, request)
	if err != nil {
		return exception.ErrInternalServer("Failed to update order")
	}

	return nil
}

func (s *ServiceImpl) Delete(ctx context.Context, ID int) error {
	orders, err := s.orderRepo.GetByIDs(ctx, []int{ID})
	if err != nil {
		return exception.ErrInternalServer("Failed to delete order")
	}

	if len(orders) < 1 {
		return exception.ErrBadRequest("Order not found")
	}

	err = s.orderRepo.Delete(ctx, ID)
	if err != nil {
		return exception.ErrInternalServer("Failed to delete order")
	}
	return nil
}

func (s *ServiceImpl) GetAll(ctx context.Context) ([]domain.Order, error) {
	orders, err := s.orderRepo.GetAll(ctx)
	if err != nil {
		return nil, exception.ErrInternalServer("Failed to delete order")

	}
	return orders, nil
}

func (s *ServiceImpl) GetByIDs(ctx context.Context, IDs []int) ([]domain.Order, error) {
	return nil, nil
}
