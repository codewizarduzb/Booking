package postgresql

import (
	"Booking/booking-service-booking/internal/entity"
	"Booking/booking-service-booking/internal/pkg/config"
	"Booking/booking-service-booking/internal/pkg/postgres"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserPostgres(t *testing.T) {
	// Connect to database
	cfg := config.New()
	db, err := postgres.New(cfg)
	if err != nil {
		return
	}

	// Test  Method Create
	repo := NewBookingRepo(db)
	hotel := &entity.GeneralBooking{
		Id:             uuid.MustParse(uuid.NewString()),
		UserId:         uuid.NewString(),
		HraId:          uuid.NewString(),
		WillArrive:     "2006-01-02",
		WillLeave:      "2006-01-03",
		NumberOfPeople: 1,
		IsCanceled:     false,
		Reason:         "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Time{},
		DeletedAt:      time.Time{},
	}

	hotelMap := make(map[string]interface{})
	hotelMap["id"] = hotel.Id.String()
	hotelMap["user_id"] = hotel.UserId
	hotelMap["hra_id"] = hotel.HraId
	hotelMap["will_arrive"] = hotel.WillArrive
	hotelMap["will_leave"] = hotel.WillLeave
	hotelMap["number_of_people"] = hotel.NumberOfPeople
	hotelMap["is_canceled"] = hotel.IsCanceled
	hotelMap["reason"] = hotel.Reason
	hotelMap["created_at"] = hotel.CreatedAt.Format("2006-01-02T15:04:05")
	// hotelMap["updated_at"] = hotel.UpdatedAt.Format("2006-01-02T15:04:05")
	// hotelMap["deleted_at"] = hotel.DeletedAt.Format("2006-01-02T15:04:05")

	ctx := context.Background()

	createdHotel, err := repo.UHBCreate(ctx, hotel)
	assert.NoError(t, err)
	assert.Equal(t, hotel.Id, createdHotel.Id)
	assert.Equal(t, hotel.UserId, createdHotel.UserId)
	assert.Equal(t, hotel.HraId, createdHotel.HraId)
	assert.Equal(t, hotel.WillArrive, createdHotel.WillArrive)
	assert.Equal(t, hotel.WillLeave, createdHotel.WillLeave)
	assert.Equal(t, hotel.NumberOfPeople, createdHotel.NumberOfPeople)
	assert.Equal(t, hotel.IsCanceled, createdHotel.IsCanceled)
	assert.Equal(t, hotel.Reason, createdHotel.Reason)
	assert.Equal(t, hotel.CreatedAt, createdHotel.CreatedAt)

	// Test Method Update
	hotel.WillArrive = "2013-01-01"
	hotel.WillArrive = "2013-01-02"
	hotel.NumberOfPeople = 2
	hotel.IsCanceled = true
	hotel.Reason = "Gender Equality"
	hotel.CreatedAt = time.Now()
	hotel.UpdatedAt = time.Now()
	updHotel, err := repo.UHBUpdate(ctx, hotel)
	assert.NoError(t, err)
	assert.Equal(t, hotel.Id, updHotel.Id)
	assert.Equal(t, hotel.UserId, updHotel.UserId)
	assert.Equal(t, hotel.HraId, updHotel.HraId)
	assert.Equal(t, hotel.WillArrive, updHotel.WillArrive)
	assert.Equal(t, hotel.WillLeave, updHotel.WillLeave)
	assert.Equal(t, hotel.NumberOfPeople, updHotel.NumberOfPeople)
	assert.Equal(t, hotel.IsCanceled, updHotel.IsCanceled)
	assert.Equal(t, hotel.Reason, updHotel.Reason)
	assert.Equal(t, hotel.CreatedAt, updHotel.CreatedAt)
	assert.Equal(t, hotel.UpdatedAt, updHotel.UpdatedAt)

	//Test Method Get
}