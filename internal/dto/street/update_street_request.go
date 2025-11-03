package street

type UpdateStreetRequest struct {
	Name string `json:"name" validate:"omitempty,min=2,max=100"`
}
