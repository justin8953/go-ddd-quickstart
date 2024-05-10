package kafka_connector

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaEventConsumer struct {
	topic  string
	reader *kafka.Reader
}

func (e *KafkaEventConsumer) Topic() string {
	return e.topic
}

func (e *KafkaEventConsumer) Listen(callback func(event map[string]interface{}) error) {
	for {
		m, err := e.reader.ReadMessage(context.Background())
		if err != nil {
			log.Error("failed to read message:", err)
		}
		var data map[string]interface{}
		if err := json.Unmarshal(m.Value, &data); err != nil {
			log.Error("failed to unmarshal message:", err)
		}
		err = callback(data)
		if err != nil {
			log.Warn("failed to process callback", err)
		}
	}
}

func (e *KafkaEventConsumer) Close() {
	if err := e.reader.Close(); err != nil {
		log.Error("failed to close reader:", err)
	}
}
