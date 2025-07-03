package domain

type Subject struct {
	Id     int64  `json:"id"`
	Name   string `json:"string"`
	Course byte   `json:"course"`
	Grades byte   `json:"grades"`
}
