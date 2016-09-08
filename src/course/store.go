package course

import (
	"database/sql"

	"github.com/mkappus/tag/src/store"
)

type CourseStore struct {
	*store.Store
}

func NewStore() (*CourseStore, error) {
	s, err := store.New("data/madison.db")
	if err != nil {
		return nil, err
	}
	return &CourseStore{s}, nil
}

func (cs *CourseStore) Create(c Course) error {
	_, err := cs.StmtExec(
		"insert into courses values(?,?,?)",
		c.Id, c.Instructor.Id, c.Title)

	if err != nil {
		return err
	}

	for _, s := range c.Pupils {
		// Courses can only add enrolled students
		if !cs.isEnrolled(s.PermId) {
			continue
		}
		_, err = cs.StmtExec(
			"insert into student_course values(?,?)",
			s.PermId, c.Id)
	}

	return err
}

func (cs *CourseStore) isEnrolled(permId int) bool {
	err := cs.QueryRow("select exists(select 1 from students where perm_id=? LIMIT=1)", permId)
	switch {
	case err == sql.ErrNoRows:
		return false

	case err != nil:
		return false

	}
	return true
}
func (cs *CourseStore) Read(id int) (Course, error) {
	var c Course

	err := cs.QueryRow("select * from courses where id=?", id).Scan(&c.Id, &c.Instructor.Id, &c.Title)
	if err != nil {
		return Course{}, err
	}

	c.Pupils, err = cs.getStudents(c.Id)
	return c, err
}

func (cs *CourseStore) getStudents(cId int) (Students, error) {
	var ss Students
	rows, err := cs.Query("select perm_id from student_course where course_id=?")
	if err != nil {
		return ss, err
	}
	for rows.Next() {
		var s Student
		_ = rows.Scan(&s.PermId)
		err := cs.QueryRow("select * from students where perm_id=?", s.PermId).Scan(&s.PermId, &s.Grade, &s.Gu, &s.FirstName, &s.LastName)
		if err != nil {
			ss = append(ss, &s)
		}
	}
	return ss, rows.Err()
}

func (cs *CourseStore) updateStudentCourse(cId, permId int) error {
	_, err := cs.StmtExec("update student_course set course_id=? where perm_id=?", cId, permId)
	return err

}

func (cs *CourseStore) Update(id int, c Course) error {
	_, err := cs.StmtExec(
		"update courses set ?, ? where id=?",
		c.Instructor.Id, c.Title, id)

	return err
}

func (cs *CourseStore) Delete(id int) error {
	_, err := cs.StmtExec(
		"delete from courses where id=?",
		id)
	return err
}
