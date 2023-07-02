package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type ParseOptions struct {
	ParseBody   bool
	ParseQuery  bool
	ParseParams bool
}

func ParseAndValidateRequest[T any](c *fiber.Ctx, request *T, options ParseOptions) error {
	// reject if request is not pointer
	value := reflect.ValueOf(request)
	if value.Kind() != reflect.Pointer {
		return fiber.NewError(fiber.StatusInternalServerError, "request must be a pointer")
	}

	// parse body
	if options.ParseBody {
		if err := c.BodyParser(request); err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
		}
	}

	// parse query
	if options.ParseQuery {
		if err := c.QueryParser(request); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "can not parse query")
		}
	}

	// parse params
	if options.ParseParams {
		if err := c.ParamsParser(request); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "can not parse params")
		}
	}

	// validate
	if err := validate.Struct(request); err != nil {
		c.Locals("request", *request)
		return err
	}

	return nil
}
