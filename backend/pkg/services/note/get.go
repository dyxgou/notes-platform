package note

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) Get(gradeId, studentId int64) (*domain.Note, *domain.AppError) {
	note, err := s.Repo.Get(gradeId, studentId)

	return note, domain.MatchError(err, "invalid gradeId or studentId")
}
