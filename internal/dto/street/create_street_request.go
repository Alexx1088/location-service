package street

type CreateStreetRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}
