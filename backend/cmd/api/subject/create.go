package subject

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/dyxgou/notas/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	s := new(core.CreateSubjectParams)

	if err := c.BodyParser(s); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := h.Validate.Struct(s); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	subject := &domain.Subject{
		Name:   s.Name,
		Course: s.Course,
		Grades: 1,
	}

	if err := h.SubjectService.Create(subject); err != nil {
		return fiber.NewError(err.Status, err.Error())
	}

	return c.JSON(fiber.Map{
		"msg": "subject created successfully",
	})
}
