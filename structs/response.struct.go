package payload_struct

type SRespPayload struct {
	Message string                 `json:"message"`
	Body    map[string]interface{} `json:"body,omitempty"`
	Detail  string                 `json:"detail,omitempty"`
}
