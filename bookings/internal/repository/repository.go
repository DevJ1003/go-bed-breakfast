package repository

import "github.com/devj1003/bookings/internal/models"

// here, we just have to add the function description
// without description it'll not work
type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
}
