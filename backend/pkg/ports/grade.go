package ports

import "github.com/dyxgou/notas/pkg/domain"

type GradeService interface {
	Create(grade *domain.Grade) (int64, *domain.AppError)
	Get(subjectId int64, name string) (*domain.Grade, *domain.AppError)
	ChangeName(id int64, name string) *domain.AppError
}

type GradeRespository interface {
	Insert(grade *domain.Grade) (int64, error)
	Get(subjectId int64, name string) (*domain.Grade, error)
	ChangeName(id int64, name string) error
}
