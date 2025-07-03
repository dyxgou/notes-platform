PRAGMA user_version = 1;

PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS student (
  id INTEGER PRIMARY KEY NOT NULL,
  name VARCHAR(30) CHECK (LENGTH (name) <= 30),
  course SMALLINT CHECK (
    course >= 0
    AND course <= 11
  ),
  parent_phone VARCHAR(10) CHECK (LENGTH (parent_phone) = 10)
);

CREATE TABLE IF NOT EXISTS subject (
  id INTEGER PRIMARY KEY NOT NULL,
  name VARCHAR(15) CHECK (LENGTH (name) <= 15),
  course SMALLINT CHECK (
    course >= 0
    AND course <= 11
  ),
  grades SMALLINT CHECK (
    grades >= 1
    AND grades <= 10
  ),
  UNIQUE (course, name)
);

CREATE TABLE IF NOT EXISTS grade (
  id INTEGER PRIMARY KEY NOT NULL,
  name VARCHAR(20) CHECK (LENGTH (name) <= 20),
  subject_id INTEGER NOT NULL,
  FOREIGN KEY (subject_id) REFERENCES subject (id)
);

CREATE TABLE IF NOT EXISTS note (
  id INTEGER PRIMARY KEY NOT NULL,
  grade_id INT NOT NULL,
  student_id INT NOT NULL,
  value SMALLINT NOT NULL CHECK (
    value >= 10
    AND value <= 50
  ),
  FOREIGN KEY (grade_id) REFERENCES grade (id),
  FOREIGN KEY (student_id) REFERENCES student (id)
);
