package repository

import (
	"time"

	"github.com/devj1003/bookings/internal/models"
)

// here, we just have to add the function description
// without description it'll not work
type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(roomID int) (string, error)
	// login, password
	GetUserByID(id int) (models.User, error)
	UpdateUser(u models.User) error
	Authenticate(email, testPassword string) (int, string, error)
	// admin
	AllReservations() ([]models.Reservation, error)
}
