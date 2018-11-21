package errors

import (
	"fmt"
	"net/http"
)

// Error is create error
func Error(code int, msg string) error {
	return &CodeError{code: code, msg: msg}
}

type CodeError struct {
	code int
	msg  string
}

func (e *CodeError) Code() int {
	return e.code
}

func (e *CodeError) Message() string {
	return e.msg
}

func (e *CodeError) Error() string {
	return e.Message()
}

func NotFound(item string) error {
	return &CodeError{code: http.StatusNotFound, msg: fmt.Sprintf("%s is not found", item)}
}
