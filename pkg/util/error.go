package util

import "github.com/sun-wenming/gin-auth/pkg/e"

// Error 带有错误码 与 错误信息的错误类
type Error interface {
	error
	Code() int
}

// ErrNew returns an error that formats as the given text.
func ErrNew(code int, text string) Error {
	return &errorString{code, text}
}

// ErrNew returns an error that formats as the given text.
func ErrNewCode(code int) Error {
	return &errorString{code, e.GetMsg(code),}
}

// ErrNew returns an error that formats as the given text.
func ErrNewSql(err error) Error {
	return &errorString{e.ErrorExecSql, err.Error(),}
}

// ErrNew returns an error that formats as the given text.
func ErrNewErr(err error) Error {
	return &errorString{e.ERROR, err.Error(),}
}

// errorString is a trivial implementation of error.
type errorString struct {
	code int
	s    string
}

func (e *errorString) Error() string {
	return e.s
}

func (e *errorString) Code() int {
	return e.code
}
