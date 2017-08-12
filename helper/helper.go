package helper

import (
	"github.com/bboortz/simplecm/operatingsystem"
	"github.com/bboortz/simplecm/reflect"
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("helper")

func init() {
	logging.SetLevel(logging.INFO, "")
}

// IsUserPresent checks is the users exists on the os
func IsUserPresent(username string) bool {
	log.Debug("HELPER: " + reflect.GetFuncName())
	exitCode, _, _ := operatingsystem.ExecCommand("id " + username)

	return exitCode == 0
}

// IsGroupPresent checks if the group exists on the os
func IsGroupPresent(username string) bool {
	log.Debug("HELPER: " + reflect.GetFuncName())
	exitCode, _, _ := operatingsystem.ExecCommand("egrep --quiet ^" + username + ": /etc/group")

	return exitCode == 0
}

// FileExists checks if the file exists
func FileExists(path string) bool {
	log.Debug("HELPER: " + reflect.GetFuncName())
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
