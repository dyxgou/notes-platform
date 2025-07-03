package student

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/dyxgou/notas/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) Create(c *fiber.Ctx) error {
	s := new(core.CreateStudentParams)

	if err := c.BodyParser(s); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := h.Validate.Struct(s); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	student := &domain.Student{
		Name:        s.Name,
		Course:      s.Course,
		ParentPhone: s.ParentPhone,
	}

	if err := h.StudentService.Create(student); err != nil {
		return fiber.NewError(err.Status, err.Error())
	}

	return c.JSON(fiber.Map{
		"msg": "student created successfully",
	})
}
