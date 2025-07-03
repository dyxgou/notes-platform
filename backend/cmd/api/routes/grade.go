package routes

import (
	"github.com/dyxgou/notas/cmd/api/grade"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) RegisterGradeGroup(router fiber.Router) {
	h := grade.NewHandler(r.Db, r.Validate)

	router.Post("/", h.Create)
	router.Get("/", h.Get)
	router.Patch("/", h.ChangeName)
}
