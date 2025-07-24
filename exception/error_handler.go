package exception

import (
	"errors"
	"steradian_code_test/web"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return web.ResponseError(ctx, code, code, err.Error())
}
