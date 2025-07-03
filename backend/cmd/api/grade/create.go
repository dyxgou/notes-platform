package grade

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/dyxgou/notas/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	p := new(core.CreateGradeParams)

	if err := c.BodyParser(p); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := h.Validate.Struct(p); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	grade := &domain.Grade{
		Name:      p.Name,
		SubjectId: p.SubjectId,
	}

	if err := h.SubjectRepo.IncrementGrades(grade.SubjectId); err != nil {
		return fiber.NewError(
			http.StatusBadRequest,
			"max amount of grades has been reached",
		)
	}

	_, err := h.GradeService.Create(grade)
	if err != nil {
		return fiber.NewError(err.Status, err.Error())
	}

	return c.JSON(fiber.Map{
		"msg": "grade created successfully",
	})
}
