package kafka_service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"

	"allocated_notifications_service/db_service"
	"allocated_notifications_service/notifications_service"
)

type Event struct {
	UserID    int                    `json:"user_id"`
	EventType string                 `json:"event_type"`
	Payload   map[string]interface{} `json:"payload"`
}

type Handler struct {
	Repo   *db_service.Repository
	Mailer *notifications_service.EmailNotifier
}

func (h *Handler) handleEvent(event Event) {
	user, err := h.Repo.GetUser(event.UserID)

	if err != nil {
		log.Printf("User not found: %v", err)
		return
	}

	msg := "Event received: " + event.EventType

	for i := 0; i < 3; i++ {
		if err := h.Mailer.SendNotification(user.Email, "Notification", msg); err != nil {
			log.Printf("Email send failed: %v", err)

			h.Repo.SaveNewNotification(db_service.Notification{
				UserID:  user.ID,
				Channel: "email",
				Message: msg,
				Status:  "failed",
			})

		} else {
			h.Repo.SaveNewNotification(db_service.Notification{
				UserID:  user.ID,
				Channel: "email",
				Message: msg,
				Status:  "sent",
			})
		}
	}
}

func (h *Handler) QueueService(brokers []string, groupID string, topics []string) {
	for _, topic := range topics {
		go func(t string) {

			r := kafka.NewReader(kafka.ReaderConfig{
				Brokers: brokers,
				GroupID: groupID,
				Topic:   t,
			})

			log.Printf("Kafka consumer started: %s", t)

			for {
				m, err := r.ReadMessage(context.Background())

				if err != nil {
					log.Printf("Kafka read error (topic %s): %v", t, err)
					continue
				}

				var ev Event

				if err := json.Unmarshal(m.Value, &ev); err != nil {
					log.Printf("Invalid event in topic %s: %v", t, err)
					continue
				}

				go h.handleEvent(ev)
			}
		}(topic)
	}
}
