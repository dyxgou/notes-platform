package grade

import (
	"os"
	"testing"

	"github.com/dyxgou/notas/pkg/domain"
	"github.com/dyxgou/notas/pkg/repositories/sqlite"
	"github.com/dyxgou/notas/pkg/repositories/sqlite/subject"
)

var r Repository
var subjectRepo subject.Repository

func TestMain(m *testing.M) {
	path := os.Getenv("DB_TEST_PATH")
	db := sqlite.ConnectClient(path)
	r.Db = db
	subjectRepo.Db = db

	code := m.Run()

	db.Close()
	os.Exit(code)
}

func TestInsertGrade(t *testing.T) {
	tt := struct {
		subject *domain.Subject
		grade   *domain.Grade
	}{
		subject: &domain.Subject{
			Name:   "Math",
			Course: 4,
			Grades: 1,
		},
		grade: &domain.Grade{
			Name: "first exam",
		},
	}

	id, err := subjectRepo.Insert(tt.subject)
	if err != nil {
		t.Fatal(err)
	}

	tt.grade.SubjectId = id

	id, err = r.Insert(tt.grade)
	if err != nil {
		t.Fatal(err)
	}

	row := r.Db.QueryRow("SELECT * FROM grade WHERE id = ?", id)

	var g domain.Grade
	if err := row.Scan(&g.Id, &g.Name, &g.SubjectId); err != nil {
		t.Fatal(err)
	}

	if g.Name != tt.grade.Name {
		t.Fatalf("grade name expected=%q. got=%q", tt.grade.Name, g.Name)
	}

	if g.SubjectId != tt.grade.SubjectId {
		t.Fatalf("grade subjectId expected=%q. got=%q", tt.grade.SubjectId, tt.grade.SubjectId)
	}
}

func TestGetGrade(t *testing.T) {
	tt := struct {
		subject *domain.Subject
		grade   *domain.Grade
	}{
		subject: &domain.Subject{
			Name:   "Math",
			Course: 7,
			Grades: 1,
		},
		grade: &domain.Grade{
			Name: "first exam",
		},
	}

	id, err := subjectRepo.Insert(tt.subject)
	if err != nil {
		t.Fatal(err)
	}

	tt.grade.SubjectId = id

	id, err = r.Insert(tt.grade)
	if err != nil {
		t.Fatal(err)
	}

	g, err := r.Get(tt.grade.SubjectId, tt.grade.Name)
	if err != nil {
		t.Fatal(err)
	}

	if g.Name != tt.grade.Name {
		t.Fatalf("grade name expected=%q. got=%q", tt.grade.Name, g.Name)
	}

	if g.SubjectId != tt.grade.SubjectId {
		t.Fatalf("grade subjectId expected=%q. got=%q", tt.grade.SubjectId, g.SubjectId)
	}
}

func TestChangeName(t *testing.T) {
	tt := struct {
		subject      *domain.Subject
		grade        *domain.Grade
		expectedName string
	}{
		subject: &domain.Subject{
			Name:   "Math",
			Course: 9,
			Grades: 1,
		},
		grade: &domain.Grade{
			Name: "first exam",
		},

		expectedName: "Second exam",
	}

	id, err := subjectRepo.Insert(tt.subject)
	if err != nil {
		t.Fatal(err)
	}

	tt.grade.SubjectId = id

	id, err = r.Insert(tt.grade)
	if err != nil {
		t.Fatal(err)
	}

	if err := r.ChangeName(id, tt.expectedName); err != nil {
		t.Fatal(err)
	}

	row := r.Db.QueryRow("SELECT name FROM grade WHERE id = ?", id)

	var name string
	if err := row.Scan(&name); err != nil {
		t.Fatal(err)
	}

	if name != tt.expectedName {
		t.Fatalf("name has not changed. expected=%q. got=%q", tt.expectedName, name)
	}
}
