package subject

func (r *Repository) IncrementGrades(id int64) error {
	query := "UPDATE subject SET grades = grades + 1 WHERE id = ?"

	_, err := r.Db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
