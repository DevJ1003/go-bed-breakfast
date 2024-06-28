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
func (m *MysqlDBRepo) InsertReservation(res models.Reservation) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insertReservationQuery := `INSERT INTO reservations (first_name, last_name, email, phone, start_date, 
					end_date, room_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
					SELECT LAST_INSERT_ID()`

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
		return err
	}

	return nil
}

// InsertRoomRestriction inserts a room restriction into the database.
// func (m *MysqlDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	insertRestrictionQuery := `INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id,
// 								created_at, updated_at, restriction_id) VALUES (?, ?, ?, ?, ?, ?, ?)`

// 	_, err := m.DB.ExecContext(ctx, insertRestrictionQuery,
// 		r.StartDate,
// 		r.EndDate,
// 		r.RoomID,
// 		r.ReservationID,
// 		time.Now(),
// 		time.Now(),
// 		r.RestrictionID,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
