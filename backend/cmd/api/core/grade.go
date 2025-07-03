package core

type CreateGradeParams struct {
	Name      string `json:"name" validate:"required,min=2,max=20"`
	SubjectId int64  `json:"subject_id" validate:"required"`
}

type GetGradeByNameAndSubjectIdQuery struct {
	Name      string `query:"name"`
	SubjectId int64  `query:"subject"`
}

type ChangeNameQuery struct {
	Id   int64  `query:"id"`
	Name string `query:"name"`
}
