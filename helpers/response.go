package helpers

type ResponseObj map[string]interface{}

type Response interface {
	SuccessWithData(string, interface{}) *response
	Success(string) *response
	Error(string) *response
}

type response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type responseImpl struct {
}

func NewResponse() Response {
	return &responseImpl{}
}

func (r *responseImpl) SuccessWithData(message string, data interface{}) *response {
	return &response{
		Error:   false,
		Message: message,
		Data:    data,
	}
}

func (r *responseImpl) Success(message string) *response {
	return &response{
		Error:   false,
		Message: message,
		Data:    nil,
	}
}

func (r *responseImpl) Error(message string) *response {
	return &response{
		Error:   true,
		Message: message,
		Data:    nil,
	}
}
