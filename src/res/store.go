package res

import (
	"github.com/mkapps/tag/src/store"
)

type ReservationStore struct {
	*store.Store
}

func NewReservationStore() {
	// TODO: Move file path to flag in main

	return store.New("data/madison.db")
}

func (rs *ReservationStore) Create(r *Reservation) error {
	result, err := rs.StmtExec(
		"insert into reservations title, created, start, end, creator_id, course_id values(?,?,?,?,?,?)",
		r.Title, r.Created, r.Start, r.End, r.CreatorId, r.CourseId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	r.Id = int(id)
	return nil
}
