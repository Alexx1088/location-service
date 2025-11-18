package crossroad

type UpdateStreetRequest struct {
	StreetId *int64 `json:"street_id" validate:"omitempty,gt=0"`
	CityId   *int64 `json:"city_id" validate:"omitempty,gt=0"`
}
