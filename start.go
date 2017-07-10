package main

import (
	"github.com/op/go-logging"
	"os"
)

func programStart() {
	envDebug := os.Getenv("DEBUG")
	if envDebug != "" {
		debug = true
	}

	// trap
	handleInterrupt(2)

	// Set the backends to be used.
	logging.SetLevel(logging.INFO, "")
	logging.SetBackend(logBackendFormatterStdout)
	log.Info("### program start ###")
	detectedOS = helperDetectOS()
	log.Debug("OS: " + detectedOS)

	currentUser = NewUser().FromCurrentUser().Build().BecomeUser()
	currentUser.CheckRoot()
}
