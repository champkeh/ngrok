package util

import (
	"runtime"
	"fmt"
)

const crashMessage = `panic: %v

%s

Oh noes! ngrok crashed!

Please submit the stack trace and any relevant information to:
github.com/inconshreveable/ngrok/issues`

func MakePanicTrace(err interface{}) string {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	return fmt.Sprintf(crashMessage, err, stackBuf[:n])
}

func PanicToError(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic: %v", r)
		}
	}()
	fn()
	return
}