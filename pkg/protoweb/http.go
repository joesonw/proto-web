package protoweb

import (
	"fmt"
)

type httpError struct {
	status int
	err    error
}

func NewHTTPError(status int, err error) error {
	return &httpError{
		status: status,
		err:    err,
	}
}

func (e *httpError) Error() string {
	return fmt.Sprintf("http %d: %s", e.status, e.err.Error())
}

func (e *httpError) Unwrap() error {
	return e.err
}

func (e *httpError) HTTPStatus() int {
	return e.status
}
