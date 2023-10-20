package kafka

import (
	"fmt"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() *ckafka.Producer {
	config := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}

	p, err := ckafka.NewProducer(config)
	if err != nil {
		panic(err)
	}

	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer, deliveryChannel chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic:     &topic,
			Partition: ckafka.PartitionAny,
		},
		Value: []byte(msg),
	}

	return producer.Produce(message, deliveryChannel)
}

func DeliveryReport(deliveryChannel chan ckafka.Event) {
	for e := range deliveryChannel {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}
	}
}
