package student

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) Delete(id int64) (int64, *domain.AppError) {
	studentId, err := s.Repo.Delete(id)

	return studentId, domain.MatchError(err, "invalid student id")
}
