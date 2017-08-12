package reflect

import (
	"runtime"
)

// GetFuncName retrieves the current function name
func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
