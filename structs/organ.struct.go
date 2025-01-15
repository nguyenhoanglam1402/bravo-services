package payload_struct

type SCreateOrganPayload struct {
	Name    string `json:"name" binding:"required"`
	ZipCode string `json:"zip_code" binding:"required"`
	Country string `json:"country" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Address string `json:"address" binding:"required"`
}
