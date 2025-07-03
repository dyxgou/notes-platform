package core

type CreateSubjectParams struct {
	Name   string `json:"name" validate:"required,min=3,max=20"`
	Course byte   `json:"course" validate:"required,gte=0,lte=11"`
}

type GetByNameAndCourseQuery struct {
	Name   string `query:"name"`
	Course byte   `query:"course"`
}
