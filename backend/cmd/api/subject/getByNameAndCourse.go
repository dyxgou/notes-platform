package subject

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetByNameAndCourse(c *fiber.Ctx) error {
	q := new(core.GetByNameAndCourseQuery)

	if err := c.QueryParser(q); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	subject, err := h.SubjectService.GetByNameAndCourse(q.Name, q.Course)

	if err != nil {
		return fiber.NewError(err.Status, err.Error())
	}

	return c.JSON(subject)
}
