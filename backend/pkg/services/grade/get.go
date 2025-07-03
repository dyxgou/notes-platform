package grade

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) Get(subjectId int64, name string) (*domain.Grade, *domain.AppError) {
	grade, err := s.Repo.Get(subjectId, name)

	return grade, domain.MatchError(err, "invalid subjectId or name")
}
