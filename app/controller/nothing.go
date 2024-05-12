package controller

import (
	"github.com/giovanniussuy/gdev-go-lang/app/audit"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get a balance
// @Description description
// @Tags tag
// @Accept json
// @Produce json
// @Param pathParam path int true "pathParam"
// @Param queryParam query int false "naoObrigatorio" Format(int)
// @Param queryParam query int true "obrigatorio" Format(int)
// @Success 200 {object} model_api_1.ModelResponse1
// @Failure 400 {object} audit.AuditResponseStatus
// @Failure 500 {object} audit.AuditResponseStatus
// @Router /v1/nothing [get]
func NothingController(webContext *fiber.Ctx) error {

	response := audit.ConstructAuditResponseStatusByCode(audit.STATUS_OK)

	return webContext.Status(response.GetStatus()).JSON(response)
}
