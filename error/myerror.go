package error

import (
	"fmt"
	"runtime"
)

func ErrorToString(err error) {
	_, file, line, ok := runtime.Caller(1)
	fmt.Printf("file: %v, line number: %v, bool: %v, error: %v\n", file, line, ok, err.Error())
}
