package assert

import (
	"fmt"
	"path"
	"runtime"
)

func True(t TBInt, expression bool) {
	if !expression {
		writeError(t, 1, "Not True")
	}
}

func Equal(t TBInt, expected interface{}, actual interface{}, args ...interface{}) {
	if expected != actual {
		writeExpectedActual(t, 1, "Not Equal", expected, actual, args...)
	}
}

func MustEqual(t TBInt, expected interface{}, actual interface{}, args ...interface{}) {
	if expected != actual {
		writeExpectedActual(t, 1, "Not Equal", expected, actual, args...)
		t.FailNow()
	}
}

func writeExpectedActual(t TBInt, skip int, prefix string, expected interface{}, actual interface{}, args ...interface{}) {
	msg := fmt.Sprintf("%s:\nExpected: %#v\n  Actual: %#v", prefix, expected, actual)

	expectedT := fmt.Sprintf("%T", expected)
	actualT := fmt.Sprintf("%T", actual)
	if expectedT != actualT {
		msg += fmt.Sprintf("\n%s != %s", expectedT, actualT)
	}

	writeError(t, skip+1, msg, args...)

}

func writeError(t TBInt, skip int, msg string, args ...interface{}) {
	// TODO use args
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		panic("can not get runtime caller")
	}
	location := fmt.Sprintf("%s/%s:%d", path.Base(path.Dir(file)), path.Base(file), line)

	t.Error(fmt.Sprintf("\r%s: %s", location, msg))
}

type TBInt interface {
	Error(args ...interface{})
	FailNow()
}
