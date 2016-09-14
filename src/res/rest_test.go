package res

func TestNew(t *testing.T) {
	t.Log("When testing if New()*Reservation:")

	s := time.Date(2016, 11, 4)
	e := time.Date(2016, 11, 5)
	e := Reservation{
		Id:        500,
		Title:     "Sample",
		Created:   time.Now(),
		Modified:  nil,
		Start:     s,
		End:       e,
		CreatorId: 600,
		CourseId:  700,
		isActive:  true,
	}
	if a := New("Sample", s, e, 600, 700); *a != *e {
		log.Error("\tExpected: %s\n\tGot: %s", e, a)
	}
}
