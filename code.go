package errorcode

type ExitCoder interface {
	ExitCode() int
}

func ExitCode(err error) int {
	return New(err).ExitCode()
}
