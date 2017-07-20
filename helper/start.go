package helper

import (
	"../logger"
	"../user"
)

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
