package helper

import (
	"github.com/bboortz/simplecm/logger"
	"github.com/bboortz/simplecm/user"
)

// ProgramStart initialize the program
func ProgramStart() {
	// trap
	handleInterrupt(2)

	// program start
	logger.InitLogger()
	log.Info("### program start ###")
	var currentUser = user.NewUser().FromCurrentUser().Build().BecomeUser()
	if !currentUser.IsRoot() {
		log.Error("You must run this program as user root!")
		ProgramExit(1)
	}

}
