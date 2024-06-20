package repository

import (
	"Booking/booking-service-booking/internal/entity"
	"context"
)

type Booking interface {
	UHBCreate(ctx context.Context, bookingHotel *entity.GeneralBooking) (*entity.GeneralBooking, error)
	URBCreate(ctx context.Context, bookingRestaurant *entity.GeneralBooking) (*entity.GeneralBooking, error)
	UABCreate(ctx context.Context, bookingAttraction *entity.GeneralBooking) (*entity.GeneralBooking, error)

	UHBGetAllByUId(ctx context.Context, limit, offset uint64, user_id string) ([]*entity.GeneralBooking, int64, error)
	URBGetAllByUId(ctx context.Context, limit, offset uint64, user_id string) ([]*entity.GeneralBooking, int64, error)
	UABGetAllByUId(ctx context.Context, limit, offset uint64, user_id string) ([]*entity.GeneralBooking, int64, error)

	UHBGetAllByHId(ctx context.Context, limit, offset uint64, room_id string) ([]*entity.Id, int64, error)
	URBGetAllByRId(ctx context.Context, limit, offset uint64, restaurant_id string) ([]*entity.Id, int64, error)
	UABGetAllByAId(ctx context.Context, limit, offset uint64, attraction_id string) ([]*entity.Id, int64, error)

	UHBList(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error)
	URBList(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error)
	UABList(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error)

	UHBListDeleted(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error)
	URBListDeleted(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error)
	UABListDeleted(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error)

	UHBUpdate(ctx context.Context, bookingHotel *entity.GeneralBooking) (*entity.GeneralBooking, error)
	URBUpdate(ctx context.Context, bookingRestaurant *entity.GeneralBooking) (*entity.GeneralBooking, error)
	UABUpdate(ctx context.Context, bookingAttraction *entity.GeneralBooking) (*entity.GeneralBooking, error)

	UHBDelete(ctx context.Context, id string) error
	URBDelete(ctx context.Context, id string) error
	UABDelete(ctx context.Context, id string) error
}
