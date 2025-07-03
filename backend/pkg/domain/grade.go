package domain

type Grade struct {
	Id        int64  `json:"id"`
	SubjectId int64  `json:"subject_id"`
	Name      string `json:"name"`
}
