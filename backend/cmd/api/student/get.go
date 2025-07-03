package student

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Get(c *fiber.Ctx) error {
	s := new(core.IdParam)

	if err := c.ParamsParser(s); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	student, err := h.StudentService.Get(s.Id)
	if err != nil {
		return fiber.NewError(err.Status, err.Error())
	}

	return c.JSON(student)
}
