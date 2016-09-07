package course

type (
	Student struct {
		PermId    int
		Grade     int
		Gu        string
		FirstName string
		LastName  string
	}

	Students []*Student

	Teacher struct {
		Id        int
		FirstName string
		LastName  string
		Email     string
		Password  string
	}

	Course struct {
		Id         int
		Instructor *Teacher
		Title      string
		Pupils     Students
	}
)
