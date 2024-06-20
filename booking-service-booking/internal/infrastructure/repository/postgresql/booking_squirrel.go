package postgresql

import (
	"Booking/booking-service-booking/internal/entity"
	"Booking/booking-service-booking/internal/pkg/otlp"
	"Booking/booking-service-booking/internal/pkg/postgres"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
)

const (
	bookingHotelTable      = "users_hotels_booking"
	bookingRestaurantTable = "users_restaurants_booking"
	bookingAttractionTable = "users_attractions_booking"
)

type bookingRepo struct {
	bookingHotelTable      string
	bookingRestaurantTable string
	bookingAttractionTable string
	db                     *postgres.PostgresDB
}

func NewBookingRepo(db *postgres.PostgresDB) *bookingRepo {
	return &bookingRepo{
		bookingHotelTable:      bookingHotelTable,
		bookingRestaurantTable: bookingRestaurantTable,
		bookingAttractionTable: bookingAttractionTable,
		db:                     db,
	}
}

// UserIdSelecter for each of tables
func (p *bookingRepo) UserIdSelecter(tableName string) squirrel.SelectBuilder {
	return p.db.Sq.Builder.Select(
		"user_id",
		"created_at",
	).From(tableName)
}

// Count for each of tables
func (p *bookingRepo) Count(tableName string) squirrel.SelectBuilder {
	return p.db.Sq.Builder.Select(
		"COUNT(*) AS count",
	).From(tableName)
}

// Selecter for each of tables
func (p *bookingRepo) Selecter(tableName string) squirrel.SelectBuilder {
	return p.db.Sq.Builder.Select(
		"id",
		"user_id",
		"hra_id",
		"TO_CHAR(will_arrive, 'YYYY-MM-DD') AS will_arrive",
		"TO_CHAR(will_leave, 'YYYY-MM-DD') AS will_leave",
		"number_of_people",
		"is_canceled",
		"reason",
		"created_at",
		"updated_at",
	).From(tableName)
}

// Create for each of tables
func (p *bookingRepo) UHBCreate(ctx context.Context, bookingHotel *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UHBCreate")
	defer span.End()

	var WA time.Time
	var WL time.Time
	var err error
	if bookingHotel.WillArrive != "" {
		WA, err = time.Parse("2006-01-02", bookingHotel.WillArrive)

		if err != nil {
			return nil, fmt.Errorf("failed to parse will arrive: %v", err)
		}
	}
	if bookingHotel.WillLeave != "" {
		WL, err = time.Parse("2006-01-02", bookingHotel.WillLeave)

		if err != nil {
			return nil, fmt.Errorf("failed to parse will leave: %v", err)
		}
	}

	data := map[string]interface{}{
		"id":               bookingHotel.Id,
		"user_id":          bookingHotel.UserId,
		"hra_id":           bookingHotel.HraId,
		"will_arrive":      WA,
		"will_leave":       WL,
		"number_of_people": bookingHotel.NumberOfPeople,
		"is_canceled":      bookingHotel.IsCanceled,
		"reason":           bookingHotel.Reason,
		"created_at":       bookingHotel.CreatedAt,
		"updated_at":       bookingHotel.UpdatedAt,
	}

	query, args, err := p.db.Sq.Builder.Insert(p.bookingHotelTable).SetMap(data).ToSql()
	if err != nil {
		return bookingHotel, fmt.Errorf("failed to build SQL query for booking hotel: %v", err)
	}
	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return bookingHotel, fmt.Errorf("failed to execute SQL query for booking hotel: %v", err)
	}

	return bookingHotel, nil
}

func (p *bookingRepo) URBCreate(ctx context.Context, bookingRestaurant *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Repository", "URBCreate")
	defer span.End()

	var WA time.Time
	var WL time.Time
	var err error
	if bookingRestaurant.WillArrive != "" {
		WA, err = time.Parse("2006-01-02", bookingRestaurant.WillArrive)

		if err != nil {
			return nil, fmt.Errorf("failed to parse will arrive: %v", err)
		}
	}
	if bookingRestaurant.WillLeave != "" {
		WL, err = time.Parse("2006-01-02", bookingRestaurant.WillLeave)

		if err != nil {
			return nil, fmt.Errorf("failed to parse will leave: %v", err)
		}
	}

	data := map[string]interface{}{
		"id":               bookingRestaurant.Id,
		"user_id":          bookingRestaurant.UserId,
		"hra_id":           bookingRestaurant.HraId,
		"will_arrive":      WA,
		"will_leave":       WL,
		"number_of_people": bookingRestaurant.NumberOfPeople,
		"is_canceled":      bookingRestaurant.IsCanceled,
		"reason":           bookingRestaurant.Reason,
		"created_at":       bookingRestaurant.CreatedAt,
		"updated_at":       bookingRestaurant.UpdatedAt,
	}

	query, args, err := p.db.Sq.Builder.Insert(p.bookingRestaurantTable).SetMap(data).ToSql()
	if err != nil {
		return bookingRestaurant, fmt.Errorf("failed to build SQL query for booking restaurant: %v", err)
	}
	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return bookingRestaurant, fmt.Errorf("failed to execute SQL query for booking restaurant: %v", err)
	}

	return bookingRestaurant, nil
}

func (p *bookingRepo) UABCreate(ctx context.Context, bookingAttraction *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UABCreate")
	defer span.End()

	var WA time.Time
	var WL time.Time
	var err error
	if bookingAttraction.WillArrive != "" {
		WA, err = time.Parse("2006-01-02", bookingAttraction.WillArrive)

		if err != nil {
			return nil, fmt.Errorf("failed to parse will arrive: %v", err)
		}
	}
	if bookingAttraction.WillLeave != "" {
		WL, err = time.Parse("2006-01-02", bookingAttraction.WillLeave)

		if err != nil {
			return nil, fmt.Errorf("failed to parse will leave: %v", err)
		}
	}

	data := map[string]interface{}{
		"id":               bookingAttraction.Id,
		"user_id":          bookingAttraction.UserId,
		"hra_id":           bookingAttraction.HraId,
		"will_arrive":      WA,
		"will_leave":       WL,
		"number_of_people": bookingAttraction.NumberOfPeople,
		"is_canceled":      bookingAttraction.IsCanceled,
		"reason":           bookingAttraction.Reason,
		"created_at":       bookingAttraction.CreatedAt,
		"updated_at":       bookingAttraction.UpdatedAt,
	}

	query, args, err := p.db.Sq.Builder.Insert(p.bookingAttractionTable).SetMap(data).ToSql()
	if err != nil {
		return bookingAttraction, fmt.Errorf("failed to build SQL query for booking attraction: %v", err)
	}
	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return bookingAttraction, fmt.Errorf("failed to execute SQL query for booking attraction: %v", err)
	}

	return bookingAttraction, nil
}

// GetAllByUId for each of tables
func (p *bookingRepo) UHBGetAllByUId(ctx context.Context, limit, offset uint64, user_id string) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UHBGetAllByUId")
	defer span.End()

	var (
		bookingHotels []*entity.GeneralBooking
		count         int64
	)

	selecter := p.Selecter(bookingHotelTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("user_id", user_id))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()
	bookingHotels = make([]*entity.GeneralBooking, 0)
	for rows.Next() {
		var bookedHotel entity.GeneralBooking
		if err = rows.Scan(
			&bookedHotel.Id,
			&bookedHotel.UserId,
			&bookedHotel.HraId,
			&bookedHotel.WillArrive,
			&bookedHotel.WillLeave,
			&bookedHotel.NumberOfPeople,
			&bookedHotel.IsCanceled,
			&bookedHotel.Reason,
			&bookedHotel.CreatedAt,
			&bookedHotel.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing bookingHotel: %v", err)
		}
		bookingHotels = append(bookingHotels, &bookedHotel)
	}

	queryCount := p.Count(bookingHotelTable)
	query, args, err = queryCount.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("user_id", user_id)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting bookingHotel: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting bookingHotel: %v", err)
	}

	return bookingHotels, count, nil
}

func (p *bookingRepo) URBGetAllByUId(ctx context.Context, limit, offset uint64, user_id string) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "URBGetAllByUId")
	defer span.End()

	var (
		bookingRestaurants []*entity.GeneralBooking
		count              int64
	)

	selecter := p.Selecter(bookingRestaurantTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("user_id", user_id))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()
	bookingRestaurants = make([]*entity.GeneralBooking, 0)
	for rows.Next() {
		var bookedRestaurant entity.GeneralBooking
		if err = rows.Scan(
			&bookedRestaurant.Id,
			&bookedRestaurant.UserId,
			&bookedRestaurant.HraId,
			&bookedRestaurant.WillArrive,
			&bookedRestaurant.WillLeave,
			&bookedRestaurant.NumberOfPeople,
			&bookedRestaurant.IsCanceled,
			&bookedRestaurant.Reason,
			&bookedRestaurant.CreatedAt,
			&bookedRestaurant.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing bookingHotel: %v", err)
		}
		bookingRestaurants = append(bookingRestaurants, &bookedRestaurant)
	}

	queryCount := p.Count(bookingRestaurantTable)
	query, args, err = queryCount.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("user_id", user_id)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting bookingHotel: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting bookingHotel: %v", err)
	}

	return bookingRestaurants, count, nil
}

func (p *bookingRepo) UABGetAllByUId(ctx context.Context, limit, offset uint64, user_id string) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UABGetAllByUId")
	defer span.End()

	var (
		bookingAttractions []*entity.GeneralBooking
		count              int64
	)

	selecter := p.Selecter(bookingAttractionTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("user_id", user_id))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()
	bookingAttractions = make([]*entity.GeneralBooking, 0)
	for rows.Next() {
		var bookedAttraction entity.GeneralBooking
		if err = rows.Scan(
			&bookedAttraction.Id,
			&bookedAttraction.UserId,
			&bookedAttraction.HraId,
			&bookedAttraction.WillArrive,
			&bookedAttraction.WillLeave,
			&bookedAttraction.NumberOfPeople,
			&bookedAttraction.IsCanceled,
			&bookedAttraction.Reason,
			&bookedAttraction.CreatedAt,
			&bookedAttraction.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing bookingHotel: %v", err)
		}
		bookingAttractions = append(bookingAttractions, &bookedAttraction)
	}

	queryCount := p.Count(bookingAttractionTable)
	query, args, err = queryCount.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("user_id", user_id)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting bookingHotel: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting bookingHotel: %v", err)
	}

	return bookingAttractions, count, nil
}

// GetAllBy HRA Id for each of tables
func (p *bookingRepo) UHBGetAllByHId(ctx context.Context, limit, offset uint64, hra_id string) ([]*entity.Id, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UHBGetAllByHId")
	defer span.End()

	var (
		userIdList []*entity.Id
		count      int64
	)

	selecter := p.UserIdSelecter(bookingHotelTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("hra_id", hra_id))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()

	userIdList = make([]*entity.Id, 0)
	for rows.Next() {
		var userId entity.Id
		if err = rows.Scan(
			&userId.UserId,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing userId: %v", err)
		}
		userIdList = append(userIdList, &userId)
	}

	queryCount := p.Count(bookingHotelTable)
	query, args, err = queryCount.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("hra_id", hra_id)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting userId: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting userId: %v", err)
	}

	return userIdList, count, nil
}

func (p *bookingRepo) URBGetAllByRId(ctx context.Context, limit, offset uint64, hra_id string) ([]*entity.Id, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "URBGetAllByRId")
	defer span.End()

	var (
		userIdList []*entity.Id
		count      int64
	)

	selecter := p.UserIdSelecter(bookingRestaurantTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("hra_id", hra_id))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()

	userIdList = make([]*entity.Id, 0)
	for rows.Next() {
		var userId entity.Id
		if err = rows.Scan(
			&userId.UserId,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing userId: %v", err)
		}
		userIdList = append(userIdList, &userId)
	}

	queryCount := p.Count(bookingRestaurantTable)
	query, args, err = queryCount.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("hra_id", hra_id)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting userId: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting userId: %v", err)
	}

	return userIdList, count, nil
}

func (p *bookingRepo) UABGetAllByAId(ctx context.Context, limit, offset uint64, hra_id string) ([]*entity.Id, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UABGetAllByAId")
	defer span.End()

	var (
		userIdList []*entity.Id
		count      int64
	)

	selecter := p.UserIdSelecter(bookingAttractionTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("hra_id", hra_id))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()

	userIdList = make([]*entity.Id, 0)
	for rows.Next() {
		var userId entity.Id
		if err = rows.Scan(
			&userId.UserId,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing userId: %v", err)
		}
		userIdList = append(userIdList, &userId)
	}

	queryCount := p.Count(bookingAttractionTable)
	query, args, err = queryCount.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.Equal("hra_id", hra_id)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting userId: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting userId: %v", err)
	}

	return userIdList, count, nil
}

// List Bookings for Admin
func (p *bookingRepo) UHBList(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UHBList")
	defer span.End()

	var (
		bookingHotels []*entity.GeneralBooking
		count         int64
	)

	selecter := p.Selecter(bookingHotelTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.Equal("deleted_at", nil))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()
	bookingHotels = make([]*entity.GeneralBooking, 0)
	for rows.Next() {
		var bookedHotel entity.GeneralBooking
		if err = rows.Scan(
			&bookedHotel.Id,
			&bookedHotel.UserId,
			&bookedHotel.HraId,
			&bookedHotel.WillArrive,
			&bookedHotel.WillLeave,
			&bookedHotel.NumberOfPeople,
			&bookedHotel.IsCanceled,
			&bookedHotel.Reason,
			&bookedHotel.CreatedAt,
			&bookedHotel.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing bookingHotel: %v", err)
		}
		bookingHotels = append(bookingHotels, &bookedHotel)
	}

	queryCount := p.Count(bookingHotelTable)
	query, args, err = queryCount.Where(p.db.Sq.Equal("deleted_at", nil)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting bookingHotel: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting bookingHotel: %v", err)
	}

	return bookingHotels, count, nil
}

func (p *bookingRepo) URBList(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "URBList")
	defer span.End()

	var (
		bookingRestaurants []*entity.GeneralBooking
		count              int64
	)

	selecter := p.Selecter(bookingRestaurantTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.Equal("deleted_at", nil))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()
	bookingRestaurants = make([]*entity.GeneralBooking, 0)
	for rows.Next() {
		var bookedRestaurant entity.GeneralBooking
		if err = rows.Scan(
			&bookedRestaurant.Id,
			&bookedRestaurant.UserId,
			&bookedRestaurant.HraId,
			&bookedRestaurant.WillArrive,
			&bookedRestaurant.WillLeave,
			&bookedRestaurant.NumberOfPeople,
			&bookedRestaurant.IsCanceled,
			&bookedRestaurant.Reason,
			&bookedRestaurant.CreatedAt,
			&bookedRestaurant.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing bookingHotel: %v", err)
		}
		bookingRestaurants = append(bookingRestaurants, &bookedRestaurant)
	}

	queryCount := p.Count(bookingRestaurantTable)
	query, args, err = queryCount.Where(p.db.Sq.Equal("deleted_at", nil)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting bookingHotel: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting bookingHotel: %v", err)
	}

	return bookingRestaurants, count, nil
}

func (p *bookingRepo) UABList(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UABList")
	defer span.End()

	var (
		bookingAttractions []*entity.GeneralBooking
		count              int64
	)

	selecter := p.Selecter(bookingAttractionTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.Equal("deleted_at", nil))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()
	bookingAttractions = make([]*entity.GeneralBooking, 0)
	for rows.Next() {
		var bookedAttraction entity.GeneralBooking
		if err = rows.Scan(
			&bookedAttraction.Id,
			&bookedAttraction.UserId,
			&bookedAttraction.HraId,
			&bookedAttraction.WillArrive,
			&bookedAttraction.WillLeave,
			&bookedAttraction.NumberOfPeople,
			&bookedAttraction.IsCanceled,
			&bookedAttraction.Reason,
			&bookedAttraction.CreatedAt,
			&bookedAttraction.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing bookingHotel: %v", err)
		}
		bookingAttractions = append(bookingAttractions, &bookedAttraction)
	}

	queryCount := p.Count(bookingAttractionTable)
	query, args, err = queryCount.Where(p.db.Sq.Equal("deleted_at", nil)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting bookingHotel: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting bookingHotel: %v", err)
	}

	return bookingAttractions, count, nil
}

// List Deleted Bookings for Admin
func (p *bookingRepo) UHBListDeleted(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UHBListDeleted")
	defer span.End()

	var (
		bookingHotels []*entity.GeneralBooking
		count         int64
	)

	selecter := p.Selecter(bookingHotelTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.NotEqual("deleted_at", nil))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()
	bookingHotels = make([]*entity.GeneralBooking, 0)
	for rows.Next() {
		var bookedHotel entity.GeneralBooking
		if err = rows.Scan(
			&bookedHotel.Id,
			&bookedHotel.UserId,
			&bookedHotel.HraId,
			&bookedHotel.WillArrive,
			&bookedHotel.WillLeave,
			&bookedHotel.NumberOfPeople,
			&bookedHotel.IsCanceled,
			&bookedHotel.Reason,
			&bookedHotel.CreatedAt,
			&bookedHotel.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing bookingHotel: %v", err)
		}
		bookingHotels = append(bookingHotels, &bookedHotel)
	}

	queryCount := p.Count(bookingHotelTable)
	query, args, err = queryCount.Where(p.db.Sq.NotEqual("deleted_at", nil)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting bookingHotel: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting bookingHotel: %v", err)
	}

	return bookingHotels, count, nil
}

func (p *bookingRepo) URBListDeleted(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "URBListDeleted")
	defer span.End()

	var (
		bookingRestaurants []*entity.GeneralBooking
		count              int64
	)

	selecter := p.Selecter(bookingRestaurantTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.NotEqual("deleted_at", nil))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()
	bookingRestaurants = make([]*entity.GeneralBooking, 0)
	for rows.Next() {
		var bookedRestaurant entity.GeneralBooking
		if err = rows.Scan(
			&bookedRestaurant.Id,
			&bookedRestaurant.UserId,
			&bookedRestaurant.HraId,
			&bookedRestaurant.WillArrive,
			&bookedRestaurant.WillLeave,
			&bookedRestaurant.NumberOfPeople,
			&bookedRestaurant.IsCanceled,
			&bookedRestaurant.Reason,
			&bookedRestaurant.CreatedAt,
			&bookedRestaurant.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing bookingHotel: %v", err)
		}
		bookingRestaurants = append(bookingRestaurants, &bookedRestaurant)
	}

	queryCount := p.Count(bookingRestaurantTable)
	query, args, err = queryCount.Where(p.db.Sq.NotEqual("deleted_at", nil)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting bookingHotel: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting bookingHotel: %v", err)
	}

	return bookingRestaurants, count, nil
}

func (p *bookingRepo) UABListDeleted(ctx context.Context, limit, offset uint64) ([]*entity.GeneralBooking, int64, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UABListDeleted")
	defer span.End()

	var (
		bookingAttractions []*entity.GeneralBooking
		count              int64
	)

	selecter := p.Selecter(bookingAttractionTable)

	if limit != 0 {
		selecter = selecter.Limit(limit).Offset(offset)
	}

	selecter = selecter.Where(p.db.Sq.NotEqual("deleted_at", nil))
	query, args, err := selecter.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()
	bookingAttractions = make([]*entity.GeneralBooking, 0)
	for rows.Next() {
		var bookedAttraction entity.GeneralBooking
		if err = rows.Scan(
			&bookedAttraction.Id,
			&bookedAttraction.UserId,
			&bookedAttraction.HraId,
			&bookedAttraction.WillArrive,
			&bookedAttraction.WillLeave,
			&bookedAttraction.NumberOfPeople,
			&bookedAttraction.IsCanceled,
			&bookedAttraction.Reason,
			&bookedAttraction.CreatedAt,
			&bookedAttraction.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing bookingHotel: %v", err)
		}
		bookingAttractions = append(bookingAttractions, &bookedAttraction)
	}

	queryCount := p.Count(bookingAttractionTable)
	query, args, err = queryCount.Where(p.db.Sq.NotEqual("deleted_at", nil)).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting bookingHotel: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row while counting bookingHotel: %v", err)
	}

	return bookingAttractions, count, nil
}

// Update
func (p *bookingRepo) UHBUpdate(ctx context.Context, bookingHotel *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UHBUpdate")
	defer span.End()

	WA, err := time.Parse("2006-01-02", bookingHotel.WillArrive)
	if err != nil {
		return nil, fmt.Errorf("failed to parse will arrive: %v", err)
	}
	WL, err := time.Parse("2006-01-02", bookingHotel.WillLeave)
	if err != nil {
		return nil, fmt.Errorf("failed to parse will leave: %v", err)
	}
	clauses := map[string]interface{}{
		"id":               bookingHotel.Id,
		"user_id":          bookingHotel.UserId,
		"hra_id":           bookingHotel.HraId,
		"will_arrive":      WA,
		"will_leave":       WL,
		"number_of_people": bookingHotel.NumberOfPeople,
		"is_canceled":      bookingHotel.IsCanceled,
		"reason":           bookingHotel.Reason,
		"created_at":       bookingHotel.CreatedAt,
		"updated_at":       bookingHotel.UpdatedAt,
	}
	sqlStr, args, err := p.db.Sq.Builder.Update(p.bookingHotelTable).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", bookingHotel.Id)).
		Where(p.db.Sq.Equal("deleted_at", nil)).
		ToSql()
	if err != nil {
		return bookingHotel, fmt.Errorf("failed to build SQL query for updating booked hotel: %v", err)
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return bookingHotel, fmt.Errorf("failed to execute SQL query for updating booked hotel: %v", err)
	}

	if commandTag.RowsAffected() == 0 {
		return bookingHotel, fmt.Errorf("no rows affected while updating booked hotel")
	}

	return bookingHotel, nil
}

func (p *bookingRepo) URBUpdate(ctx context.Context, bookingRestaurant *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Repository", "URBUpdate")
	defer span.End()

	WA, err := time.Parse("2006-01-02", bookingRestaurant.WillArrive)
	if err != nil {
		return nil, fmt.Errorf("failed to parse will arrive: %v", err)
	}
	WL, err := time.Parse("2006-01-02", bookingRestaurant.WillLeave)
	if err != nil {
		return nil, fmt.Errorf("failed to parse will leave: %v", err)
	}
	clauses := map[string]interface{}{
		"id":               bookingRestaurant.Id,
		"user_id":          bookingRestaurant.UserId,
		"hra_id":           bookingRestaurant.HraId,
		"will_arrive":      WA,
		"will_leave":       WL,
		"number_of_people": bookingRestaurant.NumberOfPeople,
		"is_canceled":      bookingRestaurant.IsCanceled,
		"reason":           bookingRestaurant.Reason,
		"created_at":       bookingRestaurant.CreatedAt,
		"updated_at":       bookingRestaurant.UpdatedAt,
	}
	sqlStr, args, err := p.db.Sq.Builder.Update(p.bookingRestaurantTable).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", bookingRestaurant.Id)).
		Where(p.db.Sq.Equal("deleted_at", nil)).
		ToSql()
	if err != nil {
		return bookingRestaurant, fmt.Errorf("failed to build SQL query for updating booked hotel: %v", err)
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return bookingRestaurant, fmt.Errorf("failed to execute SQL query for updating booked hotel: %v", err)
	}

	if commandTag.RowsAffected() == 0 {
		return bookingRestaurant, fmt.Errorf("no rows affected while updating booked hotel")
	}

	return bookingRestaurant, nil
}

func (p *bookingRepo) UABUpdate(ctx context.Context, bookingAttraction *entity.GeneralBooking) (*entity.GeneralBooking, error) {
	ctx, span := otlp.Start(ctx, "Repository", "UABUpdate")
	defer span.End()

	WA, err := time.Parse("2006-01-02", bookingAttraction.WillArrive)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date of birth: %v", err)
	}
	WL, err := time.Parse("2006-01-02", bookingAttraction.WillLeave)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date of birth: %v", err)
	}
	clauses := map[string]interface{}{
		"id":               bookingAttraction.Id,
		"user_id":          bookingAttraction.UserId,
		"hra_id":           bookingAttraction.HraId,
		"will_arrive":      WA,
		"will_leave":       WL,
		"number_of_people": bookingAttraction.NumberOfPeople,
		"is_canceled":      bookingAttraction.IsCanceled,
		"reason":           bookingAttraction.Reason,
		"created_at":       bookingAttraction.CreatedAt,
		"updated_at":       bookingAttraction.UpdatedAt,
	}
	sqlStr, args, err := p.db.Sq.Builder.Update(p.bookingAttractionTable).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", bookingAttraction.Id)).
		Where(p.db.Sq.Equal("deleted_at", nil)).
		ToSql()
	if err != nil {
		return bookingAttraction, fmt.Errorf("failed to build SQL query for updating booked hotel: %v", err)
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return bookingAttraction, fmt.Errorf("failed to execute SQL query for updating booked hotel: %v", err)
	}

	if commandTag.RowsAffected() == 0 {
		return bookingAttraction, fmt.Errorf("no rows affected while updating booked hotel")
	}

	return bookingAttraction, nil
}

// Delete
func (p *bookingRepo) UHBDelete(ctx context.Context, id string) error {
	ctx, span := otlp.Start(ctx, "Repository", "UHBDelete")
	defer span.End()

	var deletedAt sql.NullTime
	err := p.db.QueryRow(ctx, "SELECT deleted_at FROM "+p.bookingHotelTable+" WHERE id = $1", id).Scan(&deletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("%s: not found", id)
		}
		return fmt.Errorf("failed to query booking: %v", err)
	}
	if deletedAt.Valid && !deletedAt.Time.IsZero() {
		return fmt.Errorf("%s: is already soft-deleted", id)
	}

	clauses := map[string]interface{}{
		"deleted_at": time.Now().Format("2006-01-02T15:04:05"),
	}
	sqlBuilder := p.db.Sq.Builder.Update(p.bookingHotelTable).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", id))

	sqlStr, args, err := sqlBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query for canceling booking: %v", err)
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query for canceling booking: %v", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected while canceling booking")
	}

	return nil
}

func (p *bookingRepo) URBDelete(ctx context.Context, id string) error {
	ctx, span := otlp.Start(ctx, "Repository", "URBUpdate")
	defer span.End()

	var deletedAt sql.NullTime
	err := p.db.QueryRow(ctx, "SELECT deleted_at FROM "+p.bookingRestaurantTable+" WHERE id = $1", id).Scan(&deletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("%s: not found", id)
		}
		return fmt.Errorf("failed to query user: %v", err)
	}
	if deletedAt.Valid && !deletedAt.Time.IsZero() {
		return fmt.Errorf("%s: is already soft-deleted", id)
	}

	clauses := map[string]interface{}{
		"deleted_at": time.Now().Format("2006-01-02T15:04:05"),
	}
	sqlBuilder := p.db.Sq.Builder.Update(p.bookingRestaurantTable).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", id))

	sqlStr, args, err := sqlBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query for soft deleting user: %v", err)
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query for soft deleting user: %v", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected while soft deleting user")
	}

	return nil
}

func (p *bookingRepo) UABDelete(ctx context.Context, id string) error {
	ctx, span := otlp.Start(ctx, "Repository", "UABUpdate")
	defer span.End()

	var deletedAt sql.NullTime
	err := p.db.QueryRow(ctx, "SELECT deleted_at FROM "+p.bookingAttractionTable+" WHERE id = $1", id).Scan(&deletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("%s: not found", id)
		}
		return fmt.Errorf("failed to query user: %v", err)
	}
	if deletedAt.Valid && !deletedAt.Time.IsZero() {
		return fmt.Errorf("%s: is already soft-deleted", id)
	}

	clauses := map[string]interface{}{
		"deleted_at": time.Now().Format("2006-01-02T15:04:05"),
	}
	sqlBuilder := p.db.Sq.Builder.Update(p.bookingAttractionTable).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", id))

	sqlStr, args, err := sqlBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query for soft deleting user: %v", err)
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query for soft deleting user: %v", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected while soft deleting user")
	}

	return nil
}
