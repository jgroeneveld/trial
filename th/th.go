package th

import (
	"fmt"
	"path"
	"runtime"
)

func Error(t testingT, skip int, msgs ...interface{}) {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		panic("can not get runtime caller")
	}
	location := fmt.Sprintf("%s:%d", path.Base(file), line)

	args := []interface{}{fmt.Sprintf("\r%s:", location)}
	args = append(args, msgs...)
	t.Error(args...)
}

type testingT interface {
	Error(args ...interface{})
}
