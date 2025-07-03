package subject

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) Create(subject *domain.Subject) *domain.AppError {
	_, err := s.Repo.Insert(subject)

	return domain.MatchError(err, "invalid subject data")
}
