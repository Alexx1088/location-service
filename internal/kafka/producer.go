package kafka

import (
	"github.com/IBM/sarama"
)

func NewKafkaAsyncProducer(brokers []string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll

	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	return sarama.NewAsyncProducer(brokers, config)
}
