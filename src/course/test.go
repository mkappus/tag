package course

import (
	"log"
	"testing"
)

var cs *CourseStore

func init() {
	var err error
	cs, err = NewStore()
	if err != nil {
		log.Fatal(err)
	}
}

func expectError(err error, msg string, t *testing.T) {
	if err != nil {
		t.Logf("\t\tGot expected error: %s\n", err)
	} else {
		t.Errorf("\t\tDid not get expected error: %s\n", msg)
	}
}

func expectNoError(err error, msg string, t *testing.T) {
	if err == nil {
		t.Log(msg)
	} else {
		t.Errorf("Got unexpected error: %s\n", err)
	}
}

func TestIsEnrolled(t *testing.T) {
	s := &Student{PermId: 500}
	cs.Enroll(s)

	if cs.isEnrolled(s.PermId) {
		t.Log("success")
	} else {
		t.Fail()
	}
}
