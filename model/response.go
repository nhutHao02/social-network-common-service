package model

type Response[T any] struct {
	Success bool          `json:"success"`
	Data    *T            `json:"data"`
	Error   *ErrorMessage `json:"error"`
}

type PagingSuccessResponse[T any] struct {
	Success   bool `json:"success"`
	Data      *T   `json:"data"`
	TotalPage *int `json:"totalPage"`
}

type ErrorMessage struct {
	Errors  *error  `json:"errors"`
	Message *string `json:"message"`
}

func NewSuccessResponse[T any](data T) Response[T] {
	return Response[T]{
		Success: true,
		Data:    &data,
		Error:   nil,
	}
}

func NewErrorResponse(err error, message string) Response[any] {
	return Response[any]{
		Success: false,
		Data:    nil,
		Error: &ErrorMessage{
			Errors:  &err,
			Message: &message,
		},
	}
}

func NewPagingSuccessResponse[T any](data T, totalPage int) PagingSuccessResponse[T] {
	return PagingSuccessResponse[T]{
		Success:   true,
		Data:      &data,
		TotalPage: &totalPage,
	}
}
