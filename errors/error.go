package errors

import (
	"fmt"
	"github.com/mercadofarma/services/codes"
)

type Error struct {
	Code    codes.ErrCode
	Message string
}

func ErrorWithCode(code codes.ErrCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func ErrWithCode(code codes.ErrCode, err error) *Error {
	return &Error{
		Code:    code,
		Message: err.Error(),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %s - Message: %s", e.Code, e.Message)
}
