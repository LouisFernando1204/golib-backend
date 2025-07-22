package dto

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateResponseError(statusCode int, message string) Response[string] {
	return Response[string]{
		Code:    statusCode,
		Message: message,
		Data:    "",
	}
}

func CreateResponseErrorData(statusCode int, message string, data map[string]string) Response[map[string]string] {
	return Response[map[string]string]{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}
}

func CreateResponseSuccess[T any](statusCode int, data T) Response[T] {
	return Response[T]{
		Code:    statusCode,
		Message: "success",
		Data:    data,
	}
}
