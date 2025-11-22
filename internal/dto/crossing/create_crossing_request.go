package crossing

import "time"

type CreateCrossingRequest struct {
	CrossroadID int       `json:"crossroad_id" validate:"required,gt=0"`
	EventTime   time.Time `json:"event_time" validate:"required"`
}
