package repository

import (
	"time"

	"github.com/gauravpatil28/booking/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)

	InsertRoomRestriciton(r models.RoomRestriction) error

	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)

	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)

	GetRoomByID(id int) (models.Room, error)

	GetUserById(id int) (models.User, error)

	UpdateUser(u models.User) error

	Authenticate(email, testPassword string) (int, string, error)

	AllReservation() ([]models.Reservation, error)

	AllNewReservation() ([]models.Reservation, error)

	GetReservationByID(id int) (models.Reservation, error)

	UpdateReservation(u models.Reservation) error

	DeleteReservation(id int) error

	UpdateProcessedForReservation(id, processed int) error
}
