package res

type Reservation struct {
	Id        int
	Title     string
	Created   time.time
	Modified  time.Time
	Start     time.Time
	End       time.Time
	CreatorId int
	CourseId  int
	isActive  bool
}

var rs *ReservationStore

func init() {
	var err error
	rs, err = NewReservationStore()
	if err != nil {
		log.Fatal(err)
	}
}

func New(title string, start time.Time, end time.Time, creatorId, teacherId int) *Reservation {
	// TODO: validate input
	r := &Reservation{
		Title:     title,
		Created:   time.Now(),
		Start:     start,
		End:       End,
		CreatorId: creatorId,
		CourseId:  courseId,
		isActive:  true,
	}
	if err := rs.Create(r); err != nil {
		return nil
	}
}
