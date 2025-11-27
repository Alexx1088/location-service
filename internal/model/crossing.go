package model

import "time"

type Crossing struct {
	ID          int       `json:"id"`
	CrossroadID int       `json:"crossroad_id" validate:"required"`
	EventTime   time.Time `json:"event_time" validate:"required"`
}
