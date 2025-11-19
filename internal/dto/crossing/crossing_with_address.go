package crossing

import "time"

type WithAddressDTO struct {
	ID          int       `json:"id"`
	CrossroadID int       `json:"crossroad_id"`
	EventTime   time.Time `json:"event_time"`
	Street      string    `json:"street"`
	City        string    `json:"city"`
}
