package student

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Delete(c *fiber.Ctx) error {
	s := new(core.IdParam)

	if err := c.ParamsParser(s); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	id, err := h.StudentService.Delete(s.Id)
	if err != nil {
		return fiber.NewError(err.Status, err.Error())
	}

	if id == 0 {
		return fiber.NewError(http.StatusNotFound, "invalid student id")
	}

	return c.JSON(fiber.Map{
		"msg": "user deleted",
	})
}
