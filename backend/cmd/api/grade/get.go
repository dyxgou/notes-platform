package grade

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Get(c *fiber.Ctx) error {
	q := new(core.GetGradeByNameAndSubjectIdQuery)

	if err := c.QueryParser(q); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	g, err := h.GradeService.Get(q.SubjectId, q.Name)
	if err != nil {
		return fiber.NewError(err.Status, err.Error())
	}

	return c.JSON(g)
}
