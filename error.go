package errorcode

import "errors"

type Error struct {
	error
	exitCode int
}

type ErrorOpt func(*Error)

func WithExitCode(exitCode int) ErrorOpt {
	return func(e *Error) {
		e.exitCode = exitCode
	}
}

func (e *Error) ExitCode() int {
	if e == nil {
		return 0
	}

	return e.exitCode
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.error
}

func New(err error, opts ...ErrorOpt) *Error {
	e := &Error{}

	if err == nil {
		return nil
	} else if errors.As(err, &e) {
	} else {
		e = &Error{err, 1}
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}
