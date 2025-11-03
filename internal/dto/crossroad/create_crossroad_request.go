package crossroad

type CreateCrossroadRequest struct {
	StreetId int64 `json:"street_id" validate:"required,gt=0"`
	CityId   int64 `json:"city_id" validate:"required,gt=0"`
}
