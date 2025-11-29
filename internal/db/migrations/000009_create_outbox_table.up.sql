CREATE TABLE outbox (
                        id BIGSERIAL PRIMARY KEY,
                        aggregate_type TEXT NOT NULL,
                        aggregate_id BIGINT,
                        event_type TEXT NOT NULL,
                        payload JSONB NOT NULL,
                        processed BOOLEAN NOT NULL DEFAULT false,
                        processing BOOLEAN NOT NULL DEFAULT false,
                        retry_count INT NOT NULL DEFAULT 0,
                        last_error TEXT,
                        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE INDEX idx_outbox_processed ON outbox (processed, processing, created_at);