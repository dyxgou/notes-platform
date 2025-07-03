package ports

import "github.com/dyxgou/notas/pkg/domain"

type StudentService interface {
	Create(student *domain.Student) *domain.AppError
	Get(id int64) (*domain.Student, *domain.AppError)
	GetStudentsByCourse(courseId int64) ([]domain.Student, *domain.AppError)
	Delete(id int64) (int64, *domain.AppError)
}

type StudentRepository interface {
	Insert(student *domain.Student) (int64, error)
	Get(id int64) (*domain.Student, error)
	GetStudentsByCourse(courseId int64) ([]domain.Student, error)
	Delete(id int64) (int64, error)
}
