package crossing

import "time"

type FilterDTO struct {
	From        *time.Time
	To          *time.Time
	CrossroadID *int
	Limit       int
	Offset      int
}
