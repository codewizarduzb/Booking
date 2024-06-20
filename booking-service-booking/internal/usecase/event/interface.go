package event

import (
	"context"
	"Booking/booking-service-booking/internal/entity"
)

type ConsumerConfig interface {
	GetBrokers() []string
	GetTopic() string
	GetGroupID() string
	GetHandler() func(ctx context.Context, key, value []byte) error
}

type BrokerConsumer interface {
	Run()
	RegisterConsumer(config ConsumerConfig)
	Close()
}

type BrokerProducer interface {
	ProduceHotelContent(ctx context.Context, key string, value *entity.GeneralBooking) error
	ProduceRestaurantContent(ctx context.Context, key string, value *entity.GeneralBooking) error
	ProduceAttractionContent(ctx context.Context, key string, value *entity.GeneralBooking) error
	Close()
}
