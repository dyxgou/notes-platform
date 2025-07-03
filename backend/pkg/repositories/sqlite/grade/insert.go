package grade

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) Insert(g *domain.Grade) (int64, error) {
	query := "INSERT INTO grade(subject_id, name) VALUES(?, ?)"
	res, err := r.Db.Exec(query, g.SubjectId, g.Name)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
