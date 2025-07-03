package student

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) Get(id int64) (*domain.Student, *domain.AppError) {
	student, err := s.Repo.Get(id)

	if err != nil {
		return nil, domain.MatchError(err, "student not found")
	}

	return student, nil
}
