package services

import (
	pb "Booking/booking-service-booking/genproto/booking-proto"
	"Booking/booking-service-booking/internal/entity"
	"Booking/booking-service-booking/internal/pkg/otlp"
	"Booking/booking-service-booking/internal/usecase"
	"Booking/booking-service-booking/internal/usecase/event"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type bookingRPC struct {
	logger         *zap.Logger
	bookingUsecase usecase.Booking
	brokerProducer event.BrokerProducer
}

func NewRPC(logger *zap.Logger, bookingUsecase usecase.Booking, brokerProducer event.BrokerProducer) pb.BookingServiceServer {
	return &bookingRPC{
		logger:         logger,
		bookingUsecase: bookingUsecase,
		brokerProducer: brokerProducer,
	}
}

// CREATE FOR CLIENTS
func (r *bookingRPC) UHBCreate(ctx context.Context, req *pb.GeneralBook) (*pb.GeneralBook, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UHBCreate")
	span.SetAttributes(
		attribute.Key("createdId").String(req.Id),
	)
	defer span.End()

	UHBC, err := r.bookingUsecase.UHBCreate(ctx, &entity.GeneralBooking{
		Id:             uuid.Must(uuid.Parse(req.Id)),
		UserId:         req.UserId,
		HraId:          req.HraId,
		WillArrive:     req.WillArrive,
		WillLeave:      req.WillLeave,
		NumberOfPeople: req.NumberOfPeople,
		IsCanceled:     req.IsCanceled,
		Reason:         req.Reason,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GeneralBook{
		Id:             UHBC.Id.String(),
		UserId:         UHBC.UserId,
		HraId:          UHBC.HraId,
		WillArrive:     UHBC.WillArrive,
		WillLeave:      UHBC.WillLeave,
		NumberOfPeople: UHBC.NumberOfPeople,
		IsCanceled:     UHBC.IsCanceled,
		Reason:         UHBC.Reason,
		CreatedAt:      UHBC.CreatedAt.Format("2006-01-02"),
	}, nil
}

func (r *bookingRPC) URBCreate(ctx context.Context, req *pb.GeneralBook) (*pb.GeneralBook, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "URBCreate")
	span.SetAttributes(
		attribute.Key("createdId").String(req.Id),
	)
	defer span.End()

	URBC, err := r.bookingUsecase.URBCreate(ctx, &entity.GeneralBooking{
		Id:             uuid.Must(uuid.Parse(req.Id)),
		UserId:         req.UserId,
		HraId:          req.HraId,
		WillArrive:     req.WillArrive,
		NumberOfPeople: req.NumberOfPeople,
		IsCanceled:     req.IsCanceled,
		Reason:         req.Reason,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GeneralBook{
		Id:             URBC.Id.String(),
		UserId:         URBC.UserId,
		HraId:          URBC.HraId,
		WillArrive:     URBC.WillArrive,
		NumberOfPeople: URBC.NumberOfPeople,
		IsCanceled:     URBC.IsCanceled,
		Reason:         URBC.Reason,
		CreatedAt:      URBC.CreatedAt.Format("2006-01-02"),
	}, nil
}

func (r *bookingRPC) UABCreate(ctx context.Context, req *pb.GeneralBook) (*pb.GeneralBook, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UABCreate")
	span.SetAttributes(
		attribute.Key("createdId").String(req.Id),
	)
	defer span.End()

	UABC, err := r.bookingUsecase.UABCreate(ctx, &entity.GeneralBooking{
		Id:             uuid.Must(uuid.Parse(req.Id)),
		UserId:         req.UserId,
		HraId:          req.HraId,
		WillArrive:     req.WillArrive,
		NumberOfPeople: req.NumberOfPeople,
		IsCanceled:     req.IsCanceled,
		Reason:         req.Reason,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GeneralBook{
		Id:             UABC.Id.String(),
		UserId:         UABC.UserId,
		HraId:          UABC.HraId,
		WillArrive:     UABC.WillArrive,
		NumberOfPeople: UABC.NumberOfPeople,
		IsCanceled:     UABC.IsCanceled,
		Reason:         UABC.Reason,
		CreatedAt:      UABC.CreatedAt.Format("2006-01-02"),
	}, nil
}

// GET ALL BY USER ID FOR CLIENTS
func (r *bookingRPC) UHBGetAllByUId(ctx context.Context, req *pb.ListReqById) (*pb.ListUserHotelRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UHBGetAllByUId")
	span.SetAttributes(
		attribute.Key("HId").String(req.Id.Id),
		attribute.Key("Limit").Int64(int64(req.Limit)),
		attribute.Key("Offset").Int64(int64(req.Offset)),
	)
	defer span.End()

	UHBGA, count, err := r.bookingUsecase.UHBGetAllByUId(ctx, req.Limit, req.Offset, req.Id.Id)
	if err != nil {
		return nil, err
	}
	var uhbs []*pb.GeneralBook
	for _, uhb := range UHBGA {
		uhbs = append(uhbs, &pb.GeneralBook{
			Id:             uhb.Id.String(),
			UserId:         uhb.UserId,
			HraId:          uhb.HraId,
			WillArrive:     uhb.WillArrive,
			WillLeave:      uhb.WillLeave,
			NumberOfPeople: uhb.NumberOfPeople,
			IsCanceled:     uhb.IsCanceled,
			Reason:         uhb.Reason,
			CreatedAt:      uhb.CreatedAt.Format("2006-01-02"),
			UpdatedAt:      uhb.UpdatedAt.Format("2006-01-02"),
		})
	}
	return &pb.ListUserHotelRes{
		UserHotel: uhbs,
		Count:     count,
	}, nil
}

func (r *bookingRPC) URBGetAllByUId(ctx context.Context, req *pb.ListReqById) (*pb.ListUserRestaurantRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "URBGetAllByUId")
	span.SetAttributes(
		attribute.Key("RId").String(req.Id.Id),
		attribute.Key("Limit").Int64(int64(req.Limit)),
		attribute.Key("Offset").Int64(int64(req.Offset)),
	)
	defer span.End()

	URBGA, count, err := r.bookingUsecase.URBGetAllByUId(ctx, req.Limit, req.Offset, req.Id.Id)
	if err != nil {
		return nil, err
	}
	var urbs []*pb.GeneralBook
	for _, urb := range URBGA {
		urbs = append(urbs, &pb.GeneralBook{
			Id:             urb.Id.String(),
			UserId:         urb.UserId,
			HraId:          urb.HraId,
			WillArrive:     urb.WillArrive,
			NumberOfPeople: urb.NumberOfPeople,
			IsCanceled:     urb.IsCanceled,
			Reason:         urb.Reason,
			CreatedAt:      urb.CreatedAt.Format("2006-01-02"),
			UpdatedAt:      urb.UpdatedAt.Format("2006-01-02"),
		})
	}
	return &pb.ListUserRestaurantRes{
		UserRestaurant: urbs,
		Count:          count,
	}, nil
}

func (r *bookingRPC) UABGetAllByUId(ctx context.Context, req *pb.ListReqById) (*pb.ListUserAttractionRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UABGetAllByUId")
	span.SetAttributes(
		attribute.Key("AId").String(req.Id.Id),
		attribute.Key("Limit").Int64(int64(req.Limit)),
		attribute.Key("Offset").Int64(int64(req.Offset)),
	)
	defer span.End()

	UABGA, count, err := r.bookingUsecase.UABGetAllByUId(ctx, req.Limit, req.Offset, req.Id.Id)
	if err != nil {
		return nil, err
	}
	var uabs []*pb.GeneralBook
	for _, uab := range UABGA {
		uabs = append(uabs, &pb.GeneralBook{
			Id:             uab.Id.String(),
			UserId:         uab.UserId,
			HraId:          uab.HraId,
			WillArrive:     uab.WillArrive,
			NumberOfPeople: uab.NumberOfPeople,
			IsCanceled:     uab.IsCanceled,
			Reason:         uab.Reason,
			CreatedAt:      uab.CreatedAt.Format("2006-01-02"),
			UpdatedAt:      uab.UpdatedAt.Format("2006-01-02"),
		})
	}
	return &pb.ListUserAttractionRes{
		UserAttraction: uabs,
		Count:          count,
	}, nil
}

// GET ALL BY HRA ID FOR OWNERS
func (r *bookingRPC) UHBGetAllByHId(ctx context.Context, req *pb.ListReqById) (*pb.UserId, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UHBGetAllByHId")
	span.SetAttributes(
		attribute.Key("HId").String(req.Id.Id),
		attribute.Key("Limit").Int64(int64(req.Limit)),
		attribute.Key("Offset").Int64(int64(req.Offset)),
	)
	defer span.End()

	UHBGA, count, err := r.bookingUsecase.UHBGetAllByHId(ctx, req.Limit, req.Offset, req.Id.Id)
	if err != nil {
		return nil, err
	}
	var uhbs []*pb.Id
	for _, uhb := range UHBGA {
		uhbs = append(uhbs, &pb.Id{
			Id: uhb.UserId,
		})
	}
	return &pb.UserId{
		UserId: uhbs,
		Count:  count,
	}, nil
}

func (r *bookingRPC) URBGetAllByRId(ctx context.Context, req *pb.ListReqById) (*pb.UserId, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "URBGetAllByRId")
	span.SetAttributes(
		attribute.Key("RId").String(req.Id.Id),
		attribute.Key("Limit").Int64(int64(req.Limit)),
		attribute.Key("Offset").Int64(int64(req.Offset)),
	)
	defer span.End()

	URBGA, count, err := r.bookingUsecase.URBGetAllByRId(ctx, req.Limit, req.Offset, req.Id.Id)
	if err != nil {
		return nil, err
	}
	var urbs []*pb.Id
	for _, urb := range URBGA {
		urbs = append(urbs, &pb.Id{
			Id: urb.UserId,
		})
	}
	return &pb.UserId{
		UserId: urbs,
		Count:  count,
	}, nil
}

func (r *bookingRPC) UABGetAllByAId(ctx context.Context, req *pb.ListReqById) (*pb.UserId, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UABGetAllByAId")
	span.SetAttributes(
		attribute.Key("AId").String(req.Id.Id),
		attribute.Key("Limit").Int64(int64(req.Limit)),
		attribute.Key("Offset").Int64(int64(req.Offset)),
	)
	defer span.End()

	UABGA, count, err := r.bookingUsecase.UABGetAllByAId(ctx, req.Limit, req.Offset, req.Id.Id)
	if err != nil {
		return nil, err
	}
	var uabs []*pb.Id
	for _, uab := range UABGA {
		uabs = append(uabs, &pb.Id{
			Id: uab.UserId,
		})
	}
	return &pb.UserId{
		UserId: uabs,
		Count:  count,
	}, nil
}

// LIST BOOKINGS FOR ADMIN
func (r *bookingRPC) UHBList(ctx context.Context, req *pb.ListReq) (*pb.ListUserHotelRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UHBList")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(req.Limit)),
		attribute.Key("Offset").String(fmt.Sprint(req.Offset)),
	)
	defer span.End()

	println("something=======================================")
	UHBGA, count, err := r.bookingUsecase.UHBList(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	var uhbs []*pb.GeneralBook
	for _, uhb := range UHBGA {
		uhbs = append(uhbs, &pb.GeneralBook{
			Id:             uhb.Id.String(),
			UserId:         uhb.UserId,
			HraId:          uhb.HraId,
			WillArrive:     uhb.WillArrive,
			WillLeave:      uhb.WillLeave,
			NumberOfPeople: uhb.NumberOfPeople,
			IsCanceled:     uhb.IsCanceled,
			Reason:         uhb.Reason,
			CreatedAt:      uhb.CreatedAt.Format("2006-01-02"),
			UpdatedAt:      uhb.UpdatedAt.Format("2006-01-02"),
		})
	}
	return &pb.ListUserHotelRes{
		UserHotel: uhbs,
		Count:     count,
	}, nil
}

func (r *bookingRPC) URBList(ctx context.Context, req *pb.ListReq) (*pb.ListUserRestaurantRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "URBList")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(req.Limit)),
		attribute.Key("Offset").String(fmt.Sprint(req.Offset)),
	)
	defer span.End()

	URBGA, count, err := r.bookingUsecase.URBList(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	var urbs []*pb.GeneralBook
	for _, urb := range URBGA {
		urbs = append(urbs, &pb.GeneralBook{
			Id:             urb.Id.String(),
			UserId:         urb.UserId,
			HraId:          urb.HraId,
			WillArrive:     urb.WillArrive,
			NumberOfPeople: urb.NumberOfPeople,
			IsCanceled:     urb.IsCanceled,
			Reason:         urb.Reason,
			CreatedAt:      urb.CreatedAt.Format("2006-01-02"),
			UpdatedAt:      urb.UpdatedAt.Format("2006-01-02"),
		})
	}
	return &pb.ListUserRestaurantRes{
		UserRestaurant: urbs,
		Count:          count,
	}, nil
}

func (r *bookingRPC) UABList(ctx context.Context, req *pb.ListReq) (*pb.ListUserAttractionRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UABList")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(req.Limit)),
		attribute.Key("Offset").String(fmt.Sprint(req.Offset)),
	)
	defer span.End()

	UABGA, count, err := r.bookingUsecase.UABList(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	var uabs []*pb.GeneralBook
	for _, uab := range UABGA {
		uabs = append(uabs, &pb.GeneralBook{
			Id:             uab.Id.String(),
			UserId:         uab.UserId,
			HraId:          uab.HraId,
			WillArrive:     uab.WillArrive,
			NumberOfPeople: uab.NumberOfPeople,
			IsCanceled:     uab.IsCanceled,
			Reason:         uab.Reason,
			CreatedAt:      uab.CreatedAt.Format("2006-01-02"),
			UpdatedAt:      uab.UpdatedAt.Format("2006-01-02"),
		})
	}
	return &pb.ListUserAttractionRes{
		UserAttraction: uabs,
		Count:          count,
	}, nil
}

// LIST DELETED BOOKINGS FOR ADMIN
func (r *bookingRPC) UHBListDeleted(ctx context.Context, req *pb.ListReq) (*pb.ListUserHotelRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UHBListDeleted")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(req.Limit)),
		attribute.Key("Offset").String(fmt.Sprint(req.Offset)),
	)
	defer span.End()

	UHBGA, count, err := r.bookingUsecase.UHBListDeleted(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	var uhbs []*pb.GeneralBook
	for _, uhb := range UHBGA {
		uhbs = append(uhbs, &pb.GeneralBook{
			Id:             uhb.Id.String(),
			UserId:         uhb.UserId,
			HraId:          uhb.HraId,
			WillArrive:     uhb.WillArrive,
			WillLeave:      uhb.WillLeave,
			NumberOfPeople: uhb.NumberOfPeople,
			IsCanceled:     uhb.IsCanceled,
			Reason:         uhb.Reason,
			CreatedAt:      uhb.CreatedAt.Format("2006-01-02"),
			UpdatedAt:      uhb.UpdatedAt.Format("2006-01-02"),
		})
	}
	return &pb.ListUserHotelRes{
		UserHotel: uhbs,
		Count:     count,
	}, nil
}

func (r *bookingRPC) URBListDeleted(ctx context.Context, req *pb.ListReq) (*pb.ListUserRestaurantRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "URBListDeleted")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(req.Limit)),
		attribute.Key("Offset").String(fmt.Sprint(req.Offset)),
	)
	defer span.End()

	URBGA, count, err := r.bookingUsecase.URBListDeleted(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	var urbs []*pb.GeneralBook
	for _, urb := range URBGA {
		urbs = append(urbs, &pb.GeneralBook{
			Id:             urb.Id.String(),
			UserId:         urb.UserId,
			HraId:          urb.HraId,
			WillArrive:     urb.WillArrive,
			NumberOfPeople: urb.NumberOfPeople,
			IsCanceled:     urb.IsCanceled,
			Reason:         urb.Reason,
			CreatedAt:      urb.CreatedAt.Format("2006-01-02"),
			UpdatedAt:      urb.UpdatedAt.Format("2006-01-02"),
		})
	}
	return &pb.ListUserRestaurantRes{
		UserRestaurant: urbs,
		Count:          count,
	}, nil
}

func (r *bookingRPC) UABListDeleted(ctx context.Context, req *pb.ListReq) (*pb.ListUserAttractionRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UABListDeleted")
	span.SetAttributes(
		attribute.Key("Limit").String(fmt.Sprint(req.Limit)),
		attribute.Key("Offset").String(fmt.Sprint(req.Offset)),
	)
	defer span.End()

	UABGA, count, err := r.bookingUsecase.UABListDeleted(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	var uabs []*pb.GeneralBook
	for _, uab := range UABGA {
		uabs = append(uabs, &pb.GeneralBook{
			Id:             uab.Id.String(),
			UserId:         uab.UserId,
			HraId:          uab.HraId,
			WillArrive:     uab.WillArrive,
			NumberOfPeople: uab.NumberOfPeople,
			IsCanceled:     uab.IsCanceled,
			Reason:         uab.Reason,
			CreatedAt:      uab.CreatedAt.Format("2006-01-02"),
			UpdatedAt:      uab.UpdatedAt.Format("2006-01-02"),
		})
	}
	return &pb.ListUserAttractionRes{
		UserAttraction: uabs,
		Count:          count,
	}, nil
}

// UPDATE FOR CLIENTS
func (r *bookingRPC) UHBUpdate(ctx context.Context, req *pb.GeneralBook) (*pb.GeneralBook, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UHBUpdate")
	span.SetAttributes(
		attribute.Key("Id").String(fmt.Sprint(req.Id)),
	)
	defer span.End()

	UHBU, err := r.bookingUsecase.UHBUpdate(ctx, &entity.GeneralBooking{
		Id:             uuid.Must(uuid.Parse(req.Id)),
		UserId:         req.UserId,
		HraId:          req.HraId,
		WillArrive:     req.WillArrive,
		WillLeave:      req.WillLeave,
		NumberOfPeople: req.NumberOfPeople,
		IsCanceled:     req.IsCanceled,
		Reason:         req.Reason,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GeneralBook{
		Id:             UHBU.Id.String(),
		UserId:         UHBU.UserId,
		HraId:          UHBU.HraId,
		WillArrive:     UHBU.WillArrive,
		WillLeave:      UHBU.WillLeave,
		NumberOfPeople: UHBU.NumberOfPeople,
		IsCanceled:     UHBU.IsCanceled,
		Reason:         UHBU.Reason,
		CreatedAt:      UHBU.CreatedAt.Format("2006-01-02"),
	}, nil
}

func (r *bookingRPC) URBUpdate(ctx context.Context, req *pb.GeneralBook) (*pb.GeneralBook, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "URBUpdate")
	span.SetAttributes(
		attribute.Key("Id").String(fmt.Sprint(req.Id)),
	)
	defer span.End()

	URBU, err := r.bookingUsecase.URBUpdate(ctx, &entity.GeneralBooking{
		Id:             uuid.Must(uuid.Parse(req.Id)),
		UserId:         req.UserId,
		HraId:          req.HraId,
		WillArrive:     req.WillArrive,
		NumberOfPeople: req.NumberOfPeople,
		IsCanceled:     req.IsCanceled,
		Reason:         req.Reason,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GeneralBook{
		Id:             URBU.Id.String(),
		UserId:         URBU.UserId,
		HraId:          URBU.HraId,
		WillArrive:     URBU.WillArrive,
		NumberOfPeople: URBU.NumberOfPeople,
		IsCanceled:     URBU.IsCanceled,
		Reason:         URBU.Reason,
		CreatedAt:      URBU.CreatedAt.Format("2006-01-02"),
	}, nil
}

func (r *bookingRPC) UABUpdate(ctx context.Context, req *pb.GeneralBook) (*pb.GeneralBook, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UABUpdate")
	span.SetAttributes(
		attribute.Key("Id").String(fmt.Sprint(req.Id)),
	)
	defer span.End()

	UABU, err := r.bookingUsecase.UABUpdate(ctx, &entity.GeneralBooking{
		Id:             uuid.Must(uuid.Parse(req.Id)),
		UserId:         req.UserId,
		HraId:          req.HraId,
		WillArrive:     req.WillArrive,
		NumberOfPeople: req.NumberOfPeople,
		IsCanceled:     req.IsCanceled,
		Reason:         req.Reason,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GeneralBook{
		Id:             UABU.Id.String(),
		UserId:         UABU.UserId,
		HraId:          UABU.HraId,
		WillArrive:     UABU.WillArrive,
		NumberOfPeople: UABU.NumberOfPeople,
		IsCanceled:     UABU.IsCanceled,
		Reason:         UABU.Reason,
		CreatedAt:      UABU.CreatedAt.Format("2006-01-02"),
	}, nil
}

// SOFT DELETE FOR CLIENTS
func (r *bookingRPC) UHBDelete(ctx context.Context, req *pb.Id) (*pb.DelRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UHBDelete")
	span.SetAttributes(
		attribute.Key("Id").String(fmt.Sprint(req.Id)),
	)
	defer span.End()

	err := r.bookingUsecase.UHBDelete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DelRes{Result: "Successfully deleted"}, nil
}

func (r *bookingRPC) URBDelete(ctx context.Context, req *pb.Id) (*pb.DelRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "URBDelete")
	span.SetAttributes(
		attribute.Key("Id").String(fmt.Sprint(req.Id)),
	)
	defer span.End()

	err := r.bookingUsecase.URBDelete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DelRes{Result: "Successfully deleted"}, nil
}

func (r *bookingRPC) UABDelete(ctx context.Context, req *pb.Id) (*pb.DelRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UABDelete")
	span.SetAttributes(
		attribute.Key("Id").String(fmt.Sprint(req.Id)),
	)
	defer span.End()

	err := r.bookingUsecase.UABDelete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DelRes{Result: "Successfully deleted"}, nil
}
