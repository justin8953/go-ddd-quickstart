package kafka_connector

import (
	"os"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaService struct {
	Address []string
}

func NewKafkaService() *KafkaService {
	address := os.Getenv("KAFKA_BROKERS")
	if address == "" {
		address = "localhost:9092"
	}
	return &KafkaService{
		Address: strings.Split(address, ","),
	}
}

func (s *KafkaService) NewKafkaConsumer(topic string) *KafkaEventConsumer {
	return &KafkaEventConsumer{
		topic: topic,
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: s.Address,
			Topic:   topic,
			GroupID: "group_id",
		}),
	}
}

func (s *KafkaService) NewKafkaHandler(topic string) *KafkaEventHandler {
	return &KafkaEventHandler{
		topic: topic,
		writer: &kafka.Writer{
			Addr:     kafka.TCP(s.Address...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}
