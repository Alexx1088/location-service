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
	repo  repository.OutboxRepositoryInterface
	kafka sarama.SyncProducer
	topic string
}

func NewOutboxWorker(
	repo repository.OutboxRepositoryInterface,
	kafka sarama.SyncProducer,
	topic string,
) *OutboxWorker {
	return &OutboxWorker{
		repo:  repo,
		kafka: kafka,
		topic: topic,
	}
}

func (w *OutboxWorker) Start() {
	log.Println("[OUTBOX] Worker started")

	go func() {
		ticker := time.NewTicker(3 * time.Second)

		for range ticker.C {
			if err := w.processBatch(); err != nil {
				log.Println("[OUTBOX] Worker error:", err)
			}
		}
	}()
}

func (w *OutboxWorker) processBatch() error {
	ctx := context.Background()

	events, err := w.repo.GetUnprocessed(ctx, 20)
	if err != nil {
		return fmt.Errorf("faicled to load events: %w", err)
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
