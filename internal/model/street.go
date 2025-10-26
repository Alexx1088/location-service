package model

type Street struct {
	Id     int64  `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	CityId int64  `json:"city_id" db:"city_id"`
}
