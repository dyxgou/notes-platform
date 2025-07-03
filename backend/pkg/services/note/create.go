package note

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) Create(note *domain.Note) (int64, *domain.AppError) {
	id, err := s.Repo.Insert(note)

	return id, domain.MatchError(err, "invalid note data")
}
