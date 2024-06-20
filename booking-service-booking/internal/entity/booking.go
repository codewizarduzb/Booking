package entity

import (
	"time"

	"github.com/google/uuid"
)

type GeneralBooking struct {
	Id uuid.UUID
	UserId string
	HraId string
	WillArrive string
	WillLeave string
	NumberOfPeople int64
	IsCanceled bool
	Reason string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Id struct {
	UserId string
}