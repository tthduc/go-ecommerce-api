package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"strings"
	"time"
)

var (
	kafkaConsumer *kafka.Writer
)

const (
	kafkaUrl = "localhost:9092"
	topic    = "ecommerce-topic"
)

func getKafkaWriter(kafkaUrl, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaUrl),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func getKafkaReader(kafkaUrl, topic, groupId string) *kafka.Reader {
	brokers := strings.Split(kafkaUrl, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		Topic:          topic,
		GroupID:        groupId,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// Sell stock
func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{
		Message: msg,
		Type:    typeMsg,
	}

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("type"), c.Query("message"))
	body := make(map[string]interface{})
	body["message"] = s.Message
	body["type"] = s.Type

	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("key"),
		Value: jsonBody,
	}

	err := kafkaConsumer.WriteMessages(c, msg)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to send message to Kafka",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Message sent to Kafka successfully",
	})
}

func RegisterKafkaRoutes(id int) {
	kafkaGroupId := "ecommerce-group"
	reader := getKafkaReader(kafkaUrl, topic, kafkaGroupId)
	defer reader.Close()

	fmt.Println("Consumer id:", id)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error reading message:", err)
			continue
		}
		fmt.Println("msg:", string(msg.Value))

		var stockInfo StockInfo
		if err := json.Unmarshal(msg.Value, &stockInfo); err != nil {
			fmt.Println("Error unmarshalling message:", err)
			continue
		}

		fmt.Println("received message:", stockInfo.Message, "type:", stockInfo.Type)
	}
}

func main() {
	kafkaConsumer = getKafkaWriter(kafkaUrl, topic)
	defer kafkaConsumer.Close()

	r := gin.Default()
	r.POST("/kafka/action-stock", actionStock)

	// Start the Kafka consumer in a separate goroutine
	go RegisterKafkaRoutes(1)
	go RegisterKafkaRoutes(2)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
