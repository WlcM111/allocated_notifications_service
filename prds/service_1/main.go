package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	UserID    int                    `json:"user_id"`
	EventType string                 `json:"event_type"`
	Payload   map[string]interface{} `json:"payload"`
}

func main() {
	wr := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "service_1_events",
		Balancer: &kafka.LeastBytes{},
	}

	for {
		e := Event{
			UserID:    rand.Intn(5) + 1,
			EventType: "test_event",
			Payload: map[string]interface{}{
				"value": rand.Intn(1000),
			},
		}

		data, _ := json.Marshal(e)
		err := wr.WriteMessages(context.Background(),
			kafka.Message{Value: data},
		)

		if err != nil {
			log.Printf("Error: %v", err)

		} else {
			log.Printf("Sent event: %+v", e)
		}

		time.Sleep(5 * time.Second)
	}
}
