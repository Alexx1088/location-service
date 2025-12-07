package kafka

import (
	"github.com/IBM/sarama"
)

func NewKafkaProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForLocal

	config.Producer.Return.Successes = true

	config.Producer.Partitioner = sarama.NewRandomPartitioner

	return sarama.NewSyncProducer(brokers, config)
}
