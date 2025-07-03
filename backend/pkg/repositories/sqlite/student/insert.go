package student

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) Insert(student *domain.Student) (int64, error) {
	query := "INSERT INTO student(name, course, parent_phone) VALUES(?, ?, ?)"

	res, err := r.Db.Exec(query, student.Name, student.Course, student.ParentPhone)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
