package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}
	defer p.Close()

	topic := "my-topic"
	message := "Hello Kafka from Go!"

	// Produce message
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		log.Printf("Failed to produce message: %v\n", err)
	} else {
		fmt.Printf("Produced message to topic %s\n", topic)
	}

	// Wait for message delivery reports
	p.Flush(15 * 1000)
}
