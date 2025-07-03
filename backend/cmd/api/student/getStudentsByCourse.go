package student

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetStudentsByCourse(c *fiber.Ctx) error {
	cs := new(core.IdParam)

	if err := c.ParamsParser(cs); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	if cs.Id < 0 || cs.Id > 11 {
		return fiber.NewError(http.StatusBadRequest, "course should be between 0 and 11")
	}

	students, err := h.StudentService.GetStudentsByCourse(cs.Id)
	if err != nil {
		return fiber.NewError(err.Status, err.Error())
	}

	return c.JSON(students)
}
