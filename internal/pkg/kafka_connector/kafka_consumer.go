package kafka_connector

import "go-ddd-quickstart/internal/pkg/events"

type KafkaEventConsumer struct {
}

func (e *KafkaEventConsumer) Notify(event events.IEvent) {

}
