package grade

import "github.com/dyxgou/notas/pkg/domain"

func (r *Repository) Get(subjectId int64, name string) (*domain.Grade, error) {
	query := "SELECT * FROM grade WHERE subject_id = ? AND name = ?"

	row := r.Db.QueryRow(query, subjectId, name)

	var g domain.Grade
	err := row.Scan(&g.Id, &g.Name, &g.SubjectId)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
