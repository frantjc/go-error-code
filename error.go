package errorcode

import (
	"net/http"

	"github.com/frantjc/go-fn"
)

type Error struct {
	error
	exitCode, httpStatusCode int
}

type ErrorOpt func(*Error)

func WithHTTPStatusCode(statusCode int) ErrorOpt {
	return func(b *Error) {
		b.httpStatusCode = statusCode
	}
}

func WithExitCode(exitCode int) ErrorOpt {
	return func(b *Error) {
		b.exitCode = exitCode
	}
}

func (b *Error) ExitCode() int {
	if b == nil {
		return 0
	}

	return b.exitCode
}

func (b *Error) HTTPStatusCode() int {
	if b == nil {
		return http.StatusOK
	}

	return b.httpStatusCode
}

func (b *Error) Unwrap() error {
	if b == nil {
		return nil
	}

	return b.error
}

func New(err error, opts ...ErrorOpt) (b *Error) {
	if e, ok := err.(*Error); ok {
		e.error = err
		b = e
	} else {
		b = &Error{
			err,
			fn.Ternary(err == nil, 0, 1),
			fn.Ternary(err == nil, http.StatusOK, http.StatusInternalServerError),
		}
	}

	for _, opt := range opts {
		opt(b)
	}

	return b
}
