package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"location-service/internal/repository"
)

type OutboxWorker struct {
	repo      repository.OutboxRepositoryInterface
	kafka     sarama.AsyncProducer
	topic     string
	batchSize int
	interval  time.Duration
}

func NewOutboxWorker(
	repo repository.OutboxRepositoryInterface,
	kafka sarama.AsyncProducer,
	topic string,
	batchSize int,
	interval time.Duration,
) *OutboxWorker {
	return &OutboxWorker{
		repo:      repo,
		kafka:     kafka,
		topic:     topic,
		batchSize: batchSize,
		interval:  interval,
	}
}

func (w *OutboxWorker) Start(ctx context.Context) {
	log.Println("[OUTBOX] Worker started")

	go w.handleKafkaResponses(ctx)

	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("[OUTBOX] Worker stopped")
			return
		case <-ticker.C:
			if err := w.processBatch(ctx); err != nil {
				log.Println("[OUTBOX] Worker error:", err)
			}
		}
	}
}

func (w *OutboxWorker) handleKafkaResponses(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case success := <-w.kafka.Successes():
			_ = w.repo.MarkProcessed(context.Background(), success.Metadata.(int64))
		case fail := <-w.kafka.Errors():
			_ = w.repo.MarkFailed(context.Background(), fail.Msg.Metadata.(int64), fail.Err.Error())
		}
	}
}

func (w *OutboxWorker) processBatch(ctx context.Context) error {
	events, err := w.repo.LockUnprocessed(ctx, w.batchSize)
	if err != nil {
		return fmt.Errorf("failed to load events: %w", err)
	}

	if len(events) == 0 {
		return nil
	}

	for _, e := range events {
		msg := &sarama.ProducerMessage{
			Topic:    w.topic,
			Value:    sarama.ByteEncoder(e.Payload),
			Metadata: e.ID,
		}
		w.kafka.Input() <- msg
	}

	return nil
}
