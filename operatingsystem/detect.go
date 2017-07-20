package operatingsystem

import (
	"../reflect"
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
	if runtime.GOOS == "linux" {
		return true
	}

	return false
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
