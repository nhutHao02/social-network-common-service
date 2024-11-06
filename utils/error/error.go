package error

import "fmt"

type ResError struct {
	Code    *int
	Message string
}

func NewResError(code *int, message string) *ResError {
	return &ResError{
		Code:    code,
		Message: message,
	}
}

// Implement the Error method for the MyError struct
func (e *ResError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
