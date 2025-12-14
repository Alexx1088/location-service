package worker

import (
	"context"
	"fmt"
	"location-service/internal/model"
	"log"
	"time"

	"github.com/IBM/sarama"
	"location-service/internal/repository"
)

type OutboxWorker struct {
	repo      repository.OutboxRepositoryInterface
	kafka     sarama.SyncProducer
	topic     string
	batchSize int
	interval  time.Duration
}

func NewOutboxWorker(
	repo repository.OutboxRepositoryInterface,
	kafka sarama.SyncProducer,
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

	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("[OUTBOX] Worker stopped")
			return
		case <-ticker.C:
			if err := w.processBatch(); err != nil {
				log.Println("[OUTBOX] Worker error:", err)
			}
		}
	}
}

func (w *OutboxWorker) processBatch() error {
	ctx := context.Background()

	events, err := w.repo.LockUnprocessed(ctx, w.batchSize)
	if err != nil {
		return fmt.Errorf("failed to load events: %w", err)
	}

	if len(events) == 0 {
		return nil
	}

	for _, e := range events {
		if err := w.processEvent(ctx, e); err != nil {
			log.Println("[OUTBOX] processEvent error:", err)
		}
	}

	return nil
}

func (w *OutboxWorker) processEvent(ctx context.Context, e *model.OutboxEvent) error {

	msg := &sarama.ProducerMessage{
		Topic: w.topic,
		Value: sarama.ByteEncoder(e.Payload),
	}

	_, _, err := w.kafka.SendMessage(msg)
	if err != nil {
		_ = w.repo.MarkFailed(ctx, e.ID, err.Error())
		return fmt.Errorf("kafka send error: %w", err)
	}

	return w.repo.MarkProcessed(ctx, e.ID)
}
