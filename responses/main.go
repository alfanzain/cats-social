package responses

type Payload struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorPayload struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}
