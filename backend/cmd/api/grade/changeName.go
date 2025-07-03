package grade

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) ChangeName(c *fiber.Ctx) error {
	q := new(core.ChangeNameQuery)

	if err := c.BodyParser(q); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := h.GradeService.ChangeName(q.Id, q.Name); err != nil {
		return fiber.NewError(err.Status, err.Error())
	}

	return c.JSON(fiber.Map{"msg": "grade name changed successfully"})
}
