package model

import "time"

type OutboxEvent struct {
	ID            int64     `json:"id"`
	AggregateType string    `json:"aggregate_type"`
	AggregateID   *int64    `json:"aggregate_id,omitempty"`
	EventType     string    `json:"event_type"`
	Payload       []byte    `json:"payload"`
	Processed     bool      `json:"processed"`
	Processing    bool      `json:"processing"`
	RetryCount    int       `json:"retry_count"`
	LastError     *string   `json:"last_error,omitempty"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
