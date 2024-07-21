package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/devj1003/bookings/internal/models"
	"golang.org/x/crypto/bcrypt"
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

// SearchAvailabilityByDates returns true if availability exists for roomID and false if no availability
func (m *MysqlDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var numRows int

	query := `SELECT COUNT(id) FROM room_restrictions WHERE room_id = ? AND ? < end_date AND ? > start_date`

	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)

	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, err
	}

	return false, err
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (m *MysqlDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `SELECT r.id, r.room_name FROM rooms r WHERE r.id NOT IN 
			(SELECT rr.room_id FROM room_restrictions rr WHERE ? < rr.end_date AND ? > rr.start_date)`

	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {

		var room models.Room

		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}

// GetRoomByID gets a room by it's ID and returns room name
func (m *MysqlDBRepo) GetRoomByID(roomID int) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var roomname string

	query := `SELECT room_name FROM rooms WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, roomID)
	err := row.Scan(&roomname)

	if err != nil {
		return roomname, err
	}

	return roomname, err

}

// GetUserByID returns a user by id
func (m *MysqlDBRepo) GetUserByID(id int) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)
	var u models.User
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.AccessLevel,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUser updates a user in the database
func (m *MysqlDBRepo) UpdateUser(u models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `UPDATE users SET first_name=?, last_name=?, email=?, access_level=?, updated_at=?`

	_, err := m.DB.ExecContext(ctx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.AccessLevel,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

// Authenticate authenticates a user by comparing password entered and hashed password
func (m *MysqlDBRepo) Authenticate(email, testPassword string) (int, string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var id int
	var hashedPassword string

	query := `SELECT id, password FROM users WHERE email=?`
	row := m.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}

// AllReservations returns a slice of all reservations
func (m *MysqlDBRepo) AllReservations() ([]models.Reservation, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var reservations []models.Reservation

	query := `SELECT r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date,
				r.end_date, r.room_id, r.processed, r.created_at, r.updated_at, 
				rm.id, rm.room_name
				FROM reservations r
				LEFT JOIN rooms rm on (r.room_id = rm.id)
				ORDER BY r.start_date ASC`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}

	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.Processed,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Room.ID,
			&i.Room.RoomName,
		)

		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}

	if err = rows.Err(); err != nil {
		return reservations, err
	}

	return reservations, nil
}

// NewReservations returns a slice of all reservations
func (m *MysqlDBRepo) NewReservations() ([]models.Reservation, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var reservations []models.Reservation

	query := `SELECT r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date,
				r.end_date, r.room_id, r.created_at, r.updated_at,
				rm.id, rm.room_name
				FROM reservations r
				LEFT JOIN rooms rm on (r.room_id = rm.id)
				WHERE processed = 0
				ORDER BY r.start_date ASC`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}

	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Room.ID,
			&i.Room.RoomName,
		)

		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}

	if err = rows.Err(); err != nil {
		return reservations, err
	}

	return reservations, nil
}

// GetReservationByID returns one reservation by ID
func (m *MysqlDBRepo) GetReservationByID(id int) (models.Reservation, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var res models.Reservation

	query := `SELECT r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date,
				r.end_date, r.processed ,r.created_at, r.updated_at,
				rm.id, rm.room_name
				FROM reservations r
				LEFT JOIN rooms rm ON (r.room_id = rm.id)
				WHERE r.id=?`

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&res.ID,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Phone,
		&res.StartDate,
		&res.EndDate,
		&res.Processed,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.Room.ID,
		&res.Room.RoomName,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// UpdateReservation updates a reservation in the database
func (m *MysqlDBRepo) UpdateReservation(u models.Reservation) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `UPDATE reservations SET first_name=?, last_name=?, email=?,
			phone=?, updated_at=? WHERE id=?`

	_, err := m.DB.ExecContext(ctx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Phone,
		time.Now(),
		u.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteReservation deletes one reservation by id
func (m *MysqlDBRepo) DeleteReservation(id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `DELETE FROM reservations WHERE id=?`

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// UpdateProcessedForReservation updates processed for a reservation
func (m *MysqlDBRepo) UpdateProcessedForReservation(processed, id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `UPDATE reservations SET processed=? WHERE id=?`

	_, err := m.DB.ExecContext(ctx, query, processed, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *MysqlDBRepo) AllRooms() ([]models.Room, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `SELECT id, room_name, created_at, updated_at FROM rooms ORDER BY room_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return rooms, err
	}

	defer rows.Close()

	for rows.Next() {
		var rm models.Room

		err := rows.Scan(
			&rm.ID,
			&rm.RoomName,
			&rm.CreatedAt,
			&rm.UpdatedAt,
		)

		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, rm)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}
