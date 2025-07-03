package subject

import (
	"os"
	"testing"

	"github.com/dyxgou/notas/pkg/domain"
	"github.com/dyxgou/notas/pkg/repositories/sqlite"
)

var r Repository

func TestMain(m *testing.M) {
	path := os.Getenv("DB_TEST_PATH")
	db := sqlite.ConnectClient(path)
	r.Db = db

	code := m.Run()

	db.Close()
	os.Exit(code)
}

func subjectExists(t *testing.T, id int64) bool {
	query := "SELECT EXISTS(SELECT 1 FROM subject WHERE id = ?)"

	row := r.Db.QueryRow(query, id)
	var exists bool

	err := row.Scan(&exists)
	if err != nil {
		t.Fatal(err)
	}

	return exists
}

func TestInsertSubject(t *testing.T) {
	tt := &domain.Subject{
		Name:   "Math",
		Course: 11,
		Grades: 1,
	}

	id, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	if !subjectExists(t, id) {
		t.Fatalf("subject was not created. name=%q", tt.Name)
	}
}

func TestIncrementGrade(t *testing.T) {
	tt := &domain.Subject{
		Name:   "Math",
		Course: 10,
		Grades: 1,
	}

	id, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	if err := r.IncrementGrades(id); err != nil {
		t.Fatal(err)
	}

	var grades byte
	row := r.Db.QueryRow("SELECT grades FROM subject WHERE id = ?", id)

	if err := row.Scan(&grades); err != nil {
		t.Fatal(err)
	}

	want := tt.Grades + 1
	if grades != want {
		t.Fatalf("subject grades had not been incremented. want=%d. got=%d", want, grades)
	}
}

func TestGetByNameAndCourse(t *testing.T) {
	tt := &domain.Subject{
		Name:   "Social",
		Course: 2,
		Grades: 1,
	}

	_, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	subject, err := r.GetByNameAndCourse(tt.Name, tt.Course)
	if err != nil {
		t.Fatal(err)
	}

	if subject.Name != tt.Name {
		t.Fatalf("subject name expected=%q. got=%q", tt.Name, subject.Name)
	}

	if subject.Course != tt.Course {
		t.Fatalf("subject course expected=%q. got=%q", tt.Course, subject.Course)
	}

	if subject.Grades != tt.Grades {
		t.Fatalf("subject grades expected=%q. got=%q", tt.Grades, subject.Grades)
	}
}
