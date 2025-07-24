package domain

type Order struct {
	ID              int    `json:"id"`
	CarID           int    `json:"car_id"`
	OrderDate       string `json:"order_date"`
	PickupDate      string `json:"pickup_date"`
	DropOffDate     string `json:"drop_off_date"`
	PickupLocation  string `json:"pick_up_location"`
	DropOffLocation string `json:"drop_off_location"`
}
