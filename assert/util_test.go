package assert

import (
	"fmt"
	"strings"
)

type mockT struct {
	errors []string
	failedNow bool
}

func (t *mockT) Error(args ...interface{}) {
	t.errors = append(t.errors, toString(args))
}

func (t *mockT) FailNow() {
	t.failedNow = true
}

func toString(args []interface{}) string {
	var str []string

	for _, arg := range args {
		str = append(str, fmt.Sprintf("%s", arg))
	}

	return strings.Join(str, " ")
}
