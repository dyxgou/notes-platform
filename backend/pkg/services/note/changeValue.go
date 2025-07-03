package note

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) ChangeValue(id int64, value byte) *domain.AppError {
	err := s.Repo.ChangeValue(id, value)

	return domain.MatchError(err, "invalid id or value")
}
