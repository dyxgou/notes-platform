package subject

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) GetByNameAndCourse(name string, course byte) (*domain.Subject, *domain.AppError) {
	subject, err := s.Repo.GetByNameAndCourse(name, course)

	return subject, domain.MatchError(err, "invalid name or course")
}
