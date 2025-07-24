package carcontroller

import (
	"steradian_code_test/domain"
	"steradian_code_test/exception"
	carservice "steradian_code_test/service/car"
	"steradian_code_test/web"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	NewRouter(app *fiber.App)
}

type ControllerImpl struct {
	carservice carservice.Service
}

func (c *ControllerImpl) NewRouter(app *fiber.App) {
	order := app.Group("/car")
	order.Post("/", c.Create)
	order.Put("/", c.Update)
	order.Delete("/:id", c.Delete)
	order.Get("/all", c.GetAll)
}

func New(carservice carservice.Service) Controller {
	return &ControllerImpl{
		carservice: carservice,
	}
}

func (c *ControllerImpl) Create(ctx *fiber.Ctx) error {
	var request domain.Car
	err := ctx.BodyParser(&request)
	if err != nil {
		return exception.ErrorHandler(ctx, exception.ErrInternalServer("Gagal menguraikan permintaan"))
	}

	err = c.carservice.Create(ctx.Context(), request)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return web.Response(ctx, fiber.StatusCreated, nil)
}

func (c *ControllerImpl) Update(ctx *fiber.Ctx) error {
	var request domain.Car
	err := ctx.BodyParser(&request)
	if err != nil {
		// later log hcarservice
		return exception.ErrorHandler(ctx, exception.ErrInternalServer("Gagal menguraikan permintaan"))
	}

	err = c.carservice.Update(ctx.Context(), request)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return web.Response(ctx, fiber.StatusCreated, nil)
}

func (c *ControllerImpl) Delete(ctx *fiber.Ctx) error {
	request, err := ctx.ParamsInt("id")
	if err != nil {
		// later log here
		return exception.ErrorHandler(ctx, exception.ErrInternalServer("Gagal menguraikan permintaan"))
	}

	err = c.carservice.Delete(ctx.Context(), request)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return web.Response(ctx, fiber.StatusOK, nil)
}

func (c *ControllerImpl) GetAll(ctx *fiber.Ctx) error {

	response, err := c.carservice.GetAll(ctx.Context())
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return web.Response(ctx, fiber.StatusOK, response)
}
