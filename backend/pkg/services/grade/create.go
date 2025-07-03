package grade

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) Create(grade *domain.Grade) (int64, *domain.AppError) {
	id, err := s.Repo.Insert(grade)

	return id, domain.MatchError(err, "invalid grade data")
}
