package orderservice

import (
	"context"
	"errors"
	"steradian_code_test/domain"
	"steradian_code_test/exception"

	"github.com/jackc/pgx/v5"
)

func (s *ServiceImpl) ValidateOrder(ctx context.Context, request domain.Order) error {

	// validate empty payload
	if request.CarID < 1 ||
		len(request.PickupDate) < 1 || len(request.DropOffDate) < 1 ||
		len(request.PickupLocation) < 1 || len(request.DropOffLocation) < 1 {
		return exception.ErrBadRequest("All the payload must to be fulfilled")
	}

	// validate car avaibility
	car, err := s.carRepo.GetByID(ctx, request.CarID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return exception.ErrBadRequest("Car not found")
		}
		return exception.ErrInternalServer("Validation Failed")
	}

	// validate pickup or drop off date
	if request.PickupDate > request.DropOffDate {
		return exception.ErrBadRequest("Rent date is not valid")
	}

	// validate in-rent car
	orders, err := s.orderRepo.GetByCarID(ctx, car.ID)
	if err != nil {
		return exception.ErrInternalServer("Validation Failed")
	}
	for _, order := range orders {
		if order.ID != request.ID &&
			(order.PickupDate < request.PickupDate || order.DropOffDate > request.DropOffDate) {
			return exception.ErrBadRequest("The car is still in-use")
		}
	}
	return nil

}
