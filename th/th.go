package th

import (
	"fmt"
	"path"
	"runtime"
)

func Error(t TBInt, msg interface{}, skip int) {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		panic("can not get runtime caller")
	}
	location := fmt.Sprintf("%s/%s:%d", path.Base(path.Dir(file)), path.Base(file), line)

	t.Error(fmt.Sprintf("\r%s: %s", location, msg))
}

type TBInt interface {
	Error(args ...interface{})
}
