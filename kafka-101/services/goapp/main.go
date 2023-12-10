package main

import (
  "fmt"
  "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
  p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
  if err != nil {
    panic(err)
  }
  
  defer p.Close()
  
  // Delivery report handler for produced messages
  go func() {
    for e := range p.Events() {
      switch ev := e.(type) {
      case *kafka.Message:
        if ev.TopicPartition.Error != nil {
          fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
        } else {
          fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
        }
      }
    }
  }()
  
  // Produce a message
  topic := "myTopic"
  p.Produce(&kafka.Message{
    TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
    Value:          []byte("Hello Kafka!"),
  }, nil)
}