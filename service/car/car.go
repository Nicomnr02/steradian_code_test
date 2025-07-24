package carservice

import (
	"context"
	"steradian_code_test/domain"
	"steradian_code_test/exception"
	carrepository "steradian_code_test/repository/car"
)

type Service interface {
	Create(ctx context.Context, request domain.Car) error
	Update(ctx context.Context, request domain.Car) error
	Delete(ctx context.Context, ID int) error
	GetAll(ctx context.Context) ([]domain.Car, error)
}

type ServiceImpl struct {
	carRepo carrepository.Repository
}

func New(carRepo carrepository.Repository) Service {
	return &ServiceImpl{
		carRepo: carRepo,
	}
}

func (s *ServiceImpl) Create(ctx context.Context, request domain.Car) error {

	err := s.carRepo.Create(ctx, request)
	if err != nil {
		return exception.ErrInternalServer("Failed to create car")
	}

	return nil
}

func (s *ServiceImpl) Update(ctx context.Context, request domain.Car) error {

	err := s.carRepo.Update(ctx, request)
	if err != nil {
		return exception.ErrInternalServer("Failed to update car")
	}

	return nil
}

func (s *ServiceImpl) Delete(ctx context.Context, ID int) error {
	_, err := s.carRepo.GetByID(ctx, ID)
	if err != nil {
		return exception.ErrInternalServer("Failed to delete car")
	}

	err = s.carRepo.Delete(ctx, ID)
	if err != nil {
		return exception.ErrInternalServer("Failed to delete car")
	}
	return nil
}

func (s *ServiceImpl) GetAll(ctx context.Context) ([]domain.Car, error) {
	cars, err := s.carRepo.GetAll(ctx)
	if err != nil {
		return nil, exception.ErrInternalServer("Failed to delete car")

	}
	return cars, nil
}
