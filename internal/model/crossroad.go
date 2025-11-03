package model

type Crossroad struct {
	Id       int64 `json:"id" db:"id"`
	StreetId int64 `json:"street_id" db:"street_id"`
	CityId   int64 `json:"city_id" db:"city_id"`
}
