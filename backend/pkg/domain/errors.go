package domain

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-sqlite3"
)

type AppError struct {
	Msg    error
	Status int
}

func NewError(err error) *AppError {
	e := &AppError{}
	e.setStatus()

	return e
}

func (e *AppError) setStatus() {
	if errors.Is(e.Msg, sql.ErrNoRows) {
		e.Status = http.StatusNotFound
	}

	var sqlErr sqlite3.Error
	if errors.As(e.Msg, &sqlErr) {
		switch sqlErr.Code {
		case sqlite3.ErrConstraint:
			e.Status = http.StatusBadRequest
		}
	}
}

func (e *AppError) Error() string {
	return e.Msg.Error()
}

func (e *AppError) ToJSON() *fiber.Map {
	return &fiber.Map{
		"error": e.Msg.Error(),
	}
}
