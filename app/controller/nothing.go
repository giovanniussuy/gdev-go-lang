package controller

import "github.com/gofiber/fiber/v2"

func NothingController(webContext *fiber.Ctx) error {
	return webContext.Status(200).JSON("nothing")
}
