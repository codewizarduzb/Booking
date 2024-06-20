package kafka

import (
	"context"
	"Booking/booking-service-booking/internal/entity"
	"Booking/booking-service-booking/internal/pkg/config"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type producer struct {
	logger            *zap.Logger
	investmentCreated *kafka.Writer
}

func NewProducer(config *config.Config, logger *zap.Logger) *producer {
	return &producer{
		logger: logger,
		investmentCreated: &kafka.Writer{
			Addr:                   kafka.TCP(config.Kafka.Address...),
			Topic:                  config.Kafka.Topic.UserService,
			Balancer:               &kafka.Hash{},
			RequiredAcks:           kafka.RequireAll,
			AllowAutoTopicCreation: true,
			Async:                  true, // make the writer asynchronous
			Completion: func(messages []kafka.Message, err error) {
				if err != nil {
					logger.Error("kafka investmentCreated", zap.Error(err))
				}
				for _, message := range messages {
					logger.Sugar().Info(
						"kafka investmentCreated message",
						zap.Int("partition", message.Partition),
						zap.Int64("offset", message.Offset),
						zap.String("key", string(message.Key)),
						zap.String("value", string(message.Value)),
					)
				}
			},
		},
	}
}

func (p *producer) ProduceHotelContent(ctx context.Context, key string, value *entity.GeneralBooking) error {
	return nil
}

func (p *producer) ProduceRestaurantContent(ctx context.Context, key string, value *entity.GeneralBooking) error {
	return nil
}

func (p *producer) ProduceAttractionContent(ctx context.Context, key string, value *entity.GeneralBooking) error {
	return nil
}

func (p *producer) Close() {
	if err := p.investmentCreated.Close(); err != nil {
		p.logger.Error("error during close writer userCategoryCreated", zap.Error(err))
	}
}
