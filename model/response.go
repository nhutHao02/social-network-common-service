package model

type Response[T any] struct {
	Success bool    `json:"success"`
	Data    *T      `json:"data"`
	Error   *string `json:"error"`
}

type PagingSuccessResponse[T any] struct {
	Success   bool `json:"success"`
	Data      *T   `json:"data"`
	TotalPage *int `json:"totalPage"`
}

func NewSuccessResponse[T any](data T) Response[T] {
	return Response[T]{
		Success: true,
		Data:    &data,
		Error:   nil,
	}
}

func NewErrorResponse(message string) Response[any] {
	return Response[any]{
		Success: false,
		Data:    nil,
		Error:   &message,
	}
}

func NewPagingSuccessResponse[T any](data T, totalPage int) PagingSuccessResponse[T] {
	return PagingSuccessResponse[T]{
		Success:   true,
		Data:      &data,
		TotalPage: &totalPage,
	}
}
