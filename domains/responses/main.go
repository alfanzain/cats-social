package responses

type Payload struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
