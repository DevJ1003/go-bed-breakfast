package dbrepo

import (
	"context"
	"time"

	"github.com/devj1003/bookings/internal/models"
)

// here, we'll put any functions that we want to be
// available to the interface in "repository.go" file.
func (m *MysqlDBRepo) AllUsers() bool {

	return true
}

// InsertReservation inserts a reservation into the database.
func (m *MysqlDBRepo) InsertReservation(res models.Reservation) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insertReservationQuery := `INSERT INTO reservations (first_name, last_name, email, phone, start_date, 
					end_date, room_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := m.DB.ExecContext(ctx, insertReservationQuery,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return 0, err
	}

	var newID int
	lastRowQuery := `SELECT LAST_INSERT_ID() FROM reservations`
	err = m.DB.QueryRowContext(ctx, lastRowQuery).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

// InsertRoomRestriction inserts a room restriction into the database.
func (m *MysqlDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insertRestrictionQuery := `INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id,
								restriction_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := m.DB.ExecContext(ctx, insertRestrictionQuery,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		r.RestrictionID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
