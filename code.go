package errorcode

func ExitCode(err error) int {
	if err == nil {
		return 0
	}

	return New(err).ExitCode()
}

func HTTPStatusCode(err error) int {
	if err == nil {
		return 200
	}

	return New(err).HTTPStatusCode()
}
