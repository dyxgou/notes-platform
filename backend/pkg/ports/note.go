package ports

import "github.com/dyxgou/notas/pkg/domain"

type NoteService interface {
	Create(note *domain.Note) (int64, *domain.AppError)
	Get(gradeId, studentId int64) (*domain.Note, *domain.AppError)
	ChangeValue(id int64, value byte) *domain.AppError
}

type NoteRepository interface {
	Insert(note *domain.Note) (int64, error)
	Get(gradeId, studentId int64) (*domain.Note, error)
	ChangeValue(id int64, value byte) error
}
