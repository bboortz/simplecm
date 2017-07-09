package main

import (
	"os"
)



func helperUserPresent(username string) bool {
	log.Debug("HELPER: " + getFuncName())
	exitCode := execCommandWithExitCode("id " + username)
	if exitCode == 0 {
		return true
	}

	return false
}

func helperGroupPresent(username string) bool {
	log.Debug("HELPER: " + getFuncName())
	exitCode := execCommandWithExitCode("egrep --quiet ^" + username + ": /etc/group")
	if exitCode == 0 {
		return true
	}

	return false
}

func helperFileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func helperIsLinux() bool {
	log.Debug("HELPER: " + getFuncName())
	exitCode := execCommandWithExitCode("uname -a | grep Linux")
	if exitCode == 0 {
		return true
	}

	return false
}


func helperDetectOS() string {
	log.Debug("HELPER: " + getFuncName())

	if helperIsLinux() {
		if helperFileExists("/etc/arch-release") {
			return "linux-archlinux"
		}
	}

	log.Error("Unknown OS detected!");
	programExit(1)
	return "UNKNOWN"
}
