package dto

type Response[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateResponseError(message string) *Response[any] {
	return &Response[any]{
		Code:    "error",
		Message: message,
		Data:    nil,
	}
}
func CreateResponseErrorData(message string, data map[string]string) Response[map[string]string] {
	return Response[map[string]string]{
		Code:    "error",
		Message: message,
		Data:    data,
	}
}

func CreateResponseSuccess[T any](data T) *Response[T] {
	return &Response[T]{
		Code:    "success",
		Message: "success",
		Data:    data,
	}
}
