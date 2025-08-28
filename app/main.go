package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"allocated_notifications_service/db_service"
	"allocated_notifications_service/kafka_service"
	config "allocated_notifications_service/load_config"
	"allocated_notifications_service/notifications_service"
)

func main() {
	cfg := config.LoadConfig("/Users/zorinmihail/Desktop/allocated_notifications_service/config/config.yaml")

	repo := db_service.CreateNewDB(
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name,
	)

	ml := notifications_service.CreateNewEmail(cfg.Email.APIKey, cfg.Email.From)

	hd := &kafka_service.Handler{Repo: repo, Mailer: ml}
	go hd.QueueService(cfg.Kafka.Brokers, cfg.Kafka.GroupID, cfg.Kafka.Topics)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	ad := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("Server running on %s", ad)

	r.Run(ad)
}
