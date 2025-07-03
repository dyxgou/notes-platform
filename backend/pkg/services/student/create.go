package student

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) Create(student *domain.Student) *domain.AppError {
	_, err := s.Repo.Insert(student)

	return domain.MatchError(err, "invalid student data")
}
