package usecase

import (
	"Booking/booking-service-booking/internal/entity"
	"Booking/booking-service-booking/internal/infrastructure/repository"
	"Booking/booking-service-booking/internal/pkg/otlp"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
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

type BookingService struct {
	BaseUseCase
	repo       repository.Booking
	ctxTimeout time.Duration
}

func NewBookingService(ctxTimeout time.Duration, repo repository.Booking) BookingService {
	return BookingService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

// CREATE
func (s BookingService) UHBCreate(ctx context.Context, bookingHotel *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UHBCreate")
	span.SetAttributes(
		attribute.Key("Id").String(bookingHotel.Id.String()),
	)
	defer span.End()

	s.beforeRequest(nil, &bookingHotel.CreatedAt, nil, nil)
	return s.repo.UHBCreate(ctx, bookingHotel)
}

func (s BookingService) URBCreate(ctx context.Context, bookingRestaurant *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "URBCreate")
	span.SetAttributes(
		attribute.Key("Id").String(bookingRestaurant.Id.String()),
	)
	defer span.End()

	s.beforeRequest(nil, &bookingRestaurant.CreatedAt, nil, nil)
	return s.repo.URBCreate(ctx, bookingRestaurant)
}

func (s BookingService) UABCreate(ctx context.Context, bookingAttraction *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UABCreate")
	span.SetAttributes(
		attribute.Key("Id").String(bookingAttraction.Id.String()),
	)
	defer span.End()

	s.beforeRequest(nil, &bookingAttraction.CreatedAt, nil, nil)
	return s.repo.UABCreate(ctx, bookingAttraction)
}

// GET ALL BY USER ID
func (s BookingService) UHBGetAllByUId(ctx context.Context, limit, offset uint64, user_id string) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UHBGetAllByUId")
	span.SetAttributes(
		attribute.Key("UserId").String(user_id),
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.UHBGetAllByUId(ctx, limit, offset, user_id)
}

func (s BookingService) URBGetAllByUId(ctx context.Context, limit, offset uint64, user_id string) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "URBGetAllByUId")
	span.SetAttributes(
		attribute.Key("UserId").String(user_id),
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.URBGetAllByUId(ctx, limit, offset, user_id)
}

func (s BookingService) UABGetAllByUId(ctx context.Context, limit, offset uint64, user_id string) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UABGetAllByUId")
	span.SetAttributes(
		attribute.Key("UserId").String(user_id),
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.UABGetAllByUId(ctx, limit, offset, user_id)
}

// GET ALL BY HRA ID
func (s BookingService) UHBGetAllByHId(ctx context.Context, limit, offset uint64, hra_id string) ([]*entity.Id, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UHBGetAllByHId")
	span.SetAttributes(
		attribute.Key("HotelId").String(hra_id),
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.UHBGetAllByHId(ctx, limit, offset, hra_id)
}

func (s BookingService) URBGetAllByRId(ctx context.Context, limit, offset uint64, hra_id string) ([]*entity.Id, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "URBGetAllByRId")
	span.SetAttributes(
		attribute.Key("RestaurantId").String(hra_id),
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.URBGetAllByRId(ctx, limit, offset, hra_id)
}

func (s BookingService) UABGetAllByAId(ctx context.Context, limit, offset uint64, hra_id string) ([]*entity.Id, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UABGetAllByAId")
	span.SetAttributes(
		attribute.Key("AttractionId").String(hra_id),
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.UABGetAllByAId(ctx, limit, offset, hra_id)
}

// LIST BOOKINGS FOR ADMIN
func (s BookingService) UHBList(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UHBList")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.UHBList(ctx, limit, offset)
}

func (s BookingService) URBList(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "URBList")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.URBList(ctx, limit, offset)
}

func (s BookingService) UABList(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UABList")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.UABList(ctx, limit, offset)
}

// LIST DELETED BOOKINGS FOR ADMIN
func (s BookingService) UHBListDeleted(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UHBListDeleted")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.UHBListDeleted(ctx, limit, offset)
}

func (s BookingService) URBListDeleted(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "URBListDeleted")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.URBListDeleted(ctx, limit, offset)
}

func (s BookingService) UABListDeleted(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UABListDeleted")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(limit)),
		attribute.Key("Offset").String(fmt.Sprint(offset)),
	)
	defer span.End()

	return s.repo.UABListDeleted(ctx, limit, offset)
}

// UPDATE
func (s BookingService) UHBUpdate(ctx context.Context, bookingHotel *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UHBUpdate")
	span.SetAttributes(
		attribute.Key("Id").String(bookingHotel.Id.String()),
	)
	defer span.End()

	s.beforeRequest(nil, nil, &bookingHotel.UpdatedAt, nil)
	return s.repo.UHBUpdate(ctx, bookingHotel)
}

func (s BookingService) URBUpdate(ctx context.Context, bookingRestaurant *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "URBUpdate")
	span.SetAttributes(
		attribute.Key("Id").String(bookingRestaurant.Id.String()),
	)
	defer span.End()

	s.beforeRequest(nil, nil, &bookingRestaurant.UpdatedAt, nil)
	return s.repo.URBUpdate(ctx, bookingRestaurant)
}

func (s BookingService) UABUpdate(ctx context.Context, bookingAttraction *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Usecase", "UABUpdate")
	span.SetAttributes(
		attribute.Key("Id").String(bookingAttraction.Id.String()),
	)
	defer span.End()

	s.beforeRequest(nil, nil, &bookingAttraction.UpdatedAt, nil)
	return s.repo.UABUpdate(ctx, bookingAttraction)
}

// SOFT DELETE
func (s BookingService) UHBDelete(ctx context.Context, id string) error {
	ctx, span := otlp.Start(ctx, "Usecase", "UHBDelete")
	span.SetAttributes(
		attribute.Key("Id").String(id),
	)
	defer span.End()

	var bookingHotel entity.GeneralBooking
	bookingHotel.Id = uuid.MustParse(id)
	s.beforeRequest(nil, nil, nil, &bookingHotel.DeletedAt)
	return s.repo.UHBDelete(ctx, id)
}

func (s BookingService) URBDelete(ctx context.Context, id string) error {
	ctx, span := otlp.Start(ctx, "Usecase", "URBDelete")
	span.SetAttributes(
		attribute.Key("Id").String(id),
	)
	defer span.End()

	var bookingRestaurant entity.GeneralBooking
	bookingRestaurant.Id = uuid.Must(uuid.Parse(id))
	s.beforeRequest(nil, nil, nil, &bookingRestaurant.DeletedAt)
	return s.repo.URBDelete(ctx, id)
}

func (s BookingService) UABDelete(ctx context.Context, id string) error {
	ctx, span := otlp.Start(ctx, "Usecase", "UABDelete")
	span.SetAttributes(
		attribute.Key("Id").String(id),
	)
	defer span.End()

	var bookingAttraction entity.GeneralBooking
	bookingAttraction.Id = uuid.Must(uuid.Parse(id))
	s.beforeRequest(nil, nil, nil, &bookingAttraction.DeletedAt)
	return s.repo.UABDelete(ctx, id)
}
