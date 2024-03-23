package controller

import "github.com/gofiber/fiber/v2"

// @Summary Get a balance
// @Description description
// @Tags tag
// @Accept json
// @Produce json
// @Param pathParam path int true "pathParam"
// @Param queryParam query int false "naoObrigatorio" Format(int)
// @Param queryParam query int true "obrigatorio" Format(int)
// @Success 200 {object} model_api_1.ModelResponse1
// @Success 400 {object} audit.AuditResponseStatus
// @Success 500 {object} audit.AuditResponseStatus
func NothingController(webContext *fiber.Ctx) error {
	return webContext.Status(200).JSON("nothing")
}
