package kafka_connector

import (
	"context"
	"encoding/json"
	"errors"
	"go-ddd-quickstart/internal/pkg/events"
	"go-ddd-quickstart/internal/pkg/utils"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaEventHandler struct {
	topic  string
	writer *kafka.Writer
}

func (e *KafkaEventHandler) Topic() string {
	return e.topic
}

func (e *KafkaEventHandler) Notify(event events.IEvent) {
	w := e.writer
	retries := 3
	if num, error := utils.ConvertStringToInt(os.Getenv("KAFKA_RETRIES")); error == nil {
		retries = num
	}

	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		body, err := json.Marshal(event)
		if err != nil {
			log.Errorf("failed to marshal event %v", event)
		}
		message := kafka.Message{
			Key:   []byte(event.EventId()),
			Value: body,
		}
		// attempt to create topic prior to publishing the message
		err = w.WriteMessages(ctx, message)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			log.Errorf("unexpected error to write meesage %v", err)
		}
		break
	}

	if err := w.Close(); err != nil {
		log.Error("failed to close writer:", err)
	}
}
