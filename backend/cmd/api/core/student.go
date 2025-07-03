package core

type CreateStudentParams struct {
	Name        string `json:"name" validate:"required,min=2,max=30"`
	Course      byte   `json:"course" validate:"required,gte=0,lte=11"`
	ParentPhone string `json:"parent_phone" validate:"required,len=10"`
}
