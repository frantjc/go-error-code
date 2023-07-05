package errorcode

import "errors"

type Error struct {
	error
	exitCode, httpStatusCode int
}

type ErrorOpt func(*Error)

func WithExitCode(exitCode int) ErrorOpt {
	return func(e *Error) {
		e.exitCode = exitCode
	}
}

func WithHTTPStatusCode(httpStatusCode int) ErrorOpt {
	return func(e *Error) {
		e.httpStatusCode = httpStatusCode
	}
}

func (e *Error) ExitCode() int {
	if e == nil {
		return 0
	}

	return e.exitCode
}

func (e *Error) HTTPStatusCode() int {
	if e == nil {
		return 0
	}

	return e.httpStatusCode
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.error
}

func New(err error, opts ...ErrorOpt) *Error {
	e := &Error{}

	switch {
	case err == nil:
		return nil
	case errors.As(err, &e):
	default:
		e = &Error{err, 1, 500}
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}
