package main

import (
	"encoding/json"
	"log"
	carcontroller "steradian_code_test/controller/car"
	ordercontroller "steradian_code_test/controller/order"
	"steradian_code_test/db/postgres"
	"steradian_code_test/exception"
	carrepository "steradian_code_test/repository/car"
	orderrepository "steradian_code_test/repository/order"
	carservice "steradian_code_test/service/car"
	orderservice "steradian_code_test/service/order"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// DB
	DB := postgres.Init()

	// HTTP
	RunHTTP(DB)

}

func RunHTTP(DB *pgxpool.Pool) {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: exception.ErrorHandler,
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		},
	)

	carrepo := carrepository.New(DB)
	orderrepo := orderrepository.New(DB)

	carservice := carservice.New(carrepo)
	orderservice := orderservice.New(orderrepo, carrepo)

	carcontroller := carcontroller.New(carservice)
	ordercontroller := ordercontroller.New(orderservice)

	carcontroller.NewRouter(app)
	ordercontroller.NewRouter(app)

	err := app.Listen("0.0.0.0:5000")
	if err != nil {
		log.Fatal(err)
	}
}
