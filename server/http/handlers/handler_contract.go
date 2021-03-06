package handlers

import (
	"github.com/ecommerce-service/product-service/usecases"
	"github.com/gofiber/fiber/v2"
)

type HandlerContract struct {
	UseCaseContract *usecases.UseCaseContract
	App             *fiber.App
}
