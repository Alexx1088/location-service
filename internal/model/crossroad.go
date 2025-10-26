package model

type Crossroad struct {
	Id       int64 `json:"id" db:"id"`
	StreetId int64 `json:"street_id" db:"street_id"`
}
