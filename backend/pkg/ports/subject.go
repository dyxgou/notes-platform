package ports

import "github.com/dyxgou/notas/pkg/domain"

type SubjectService interface {
	Create(subject *domain.Subject) *domain.AppError
	GetByNameAndCourse(name string, course byte) (*domain.Subject, *domain.AppError)
}

type SubjectRepository interface {
	Insert(subject *domain.Subject) (int64, error)
	IncrementGrades(id int64) error
	GetByNameAndCourse(name string, course byte) (*domain.Subject, error)
}
