package helper

import (
	"../operatingsystem"
	"../reflect"
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("helper")

func init() {
	logging.SetLevel(logging.INFO, "")
}

func IsUserPresent(username string) bool {
	log.Debug("HELPER: " + reflect.GetFuncName())
	exitCode, _, _ := operatingsystem.ExecCommand("id " + username)
	if exitCode == 0 {
		return true
	}

	return false
}

func IsGroupPresent(username string) bool {
	log.Debug("HELPER: " + reflect.GetFuncName())
	exitCode, _, _ := operatingsystem.ExecCommand("egrep --quiet ^" + username + ": /etc/group")
	if exitCode == 0 {
		return true
	}

	return false
}

func FileExists(path string) bool {
	log.Debug("HELPER: " + reflect.GetFuncName())
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
