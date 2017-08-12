package operatingsystem

import (
	"github.com/bboortz/simplecm/reflect"
	"os"
	"runtime"
)

func fileExists(path string) bool {
	log.Debug("HELPER: " + reflect.GetFuncName())
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func isLinux() bool {
	log.Debug("HELPER: " + reflect.GetFuncName())

	return runtime.GOOS == "linux"
}

func detectOS() string {
	log.Debug("HELPER: " + reflect.GetFuncName())

	if isLinux() {
		if fileExists("/etc/arch-release") {
			return "linux-archlinux"
		}
	}

	return "UNKNOWN"
}
