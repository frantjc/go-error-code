package errorcode_test

import (
	"fmt"
	"testing"

	errorcode "github.com/frantjc/go-error-code"
)

func TestExitCode(t *testing.T) {
	var (
		expected = 11
		err      = errorcode.New(fmt.Errorf("test"), errorcode.WithExitCode(expected))
		actual   = errorcode.ExitCode(err)
	)

	if actual != expected {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}
