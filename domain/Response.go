package domain

type Response[T interface{}] struct {
	Data    *T   `json:"data" swaggertype:"object"`
	Success bool `json:"success"`
}

func GenerateErrorResponse(error string) *Response[string] {
	return &Response[string]{
		Success: false,
		Data:    &error,
	}
}

func GenerateSuccessResponse[T interface{}](data *T) *Response[T] {
	return &Response[T]{
		Success: true,
		Data:    data,
	}
}
