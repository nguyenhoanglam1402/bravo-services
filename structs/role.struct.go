package payload_struct

type CreateRolePayload struct {
	Name string `json:"name" binding:"required"`
}
