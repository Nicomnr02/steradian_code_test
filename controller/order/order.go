package ordercontroller

import (
	"steradian_code_test/domain"
	"steradian_code_test/exception"
	orderservice "steradian_code_test/service/order"
	"steradian_code_test/web"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	NewRouter(app *fiber.App)
}

type ControllerImpl struct {
	orderService orderservice.Service
}

func (c *ControllerImpl) NewRouter(app *fiber.App) {
	order := app.Group("/order")
	order.Post("/", c.Create)
	order.Put("/", c.Update)
	order.Delete("/:id", c.Delete)
	order.Get("/:id/item", c.GetByID)
	order.Get("/all", c.GetAll)
}

func New(orderService orderservice.Service) Controller {
	return &ControllerImpl{
		orderService: orderService,
	}
}

func (c *ControllerImpl) Create(ctx *fiber.Ctx) error {
	var request domain.Order
	err := ctx.BodyParser(&request)
	if err != nil {
		// later log here
		return exception.ErrorHandler(ctx, exception.ErrInternalServer("Gagal menguraikan permintaan"))
	}

	err = c.orderService.Create(ctx.Context(), request)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return web.Response(ctx, fiber.StatusCreated, nil)
}

func (c *ControllerImpl) Update(ctx *fiber.Ctx) error {
	var request domain.Order
	err := ctx.BodyParser(&request)
	if err != nil {
		// later log here
		return exception.ErrorHandler(ctx, exception.ErrInternalServer("Gagal menguraikan permintaan"))
	}

	err = c.orderService.Update(ctx.Context(), request)
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

	err = c.orderService.Delete(ctx.Context(), request)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return web.Response(ctx, fiber.StatusOK, nil)
}

func (c *ControllerImpl) GetAll(ctx *fiber.Ctx) error {

	response, err := c.orderService.GetAll(ctx.Context())
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return web.Response(ctx, fiber.StatusOK, response)
}

func (c *ControllerImpl) GetByID(ctx *fiber.Ctx) error {
	request, err := ctx.ParamsInt("id")
	if err != nil {
		// later log here
		return exception.ErrorHandler(ctx, exception.ErrInternalServer("Gagal menguraikan permintaan"))
	}

	response, err := c.orderService.GetByIDs(ctx.Context(), []int{request})
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return web.Response(ctx, fiber.StatusOK, response)
}
