package helper

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var globalExitCode = 0

func handleInterrupt(intrptChSize int) {
	s := make(chan os.Signal, intrptChSize)
	signal.Notify(s,
		syscall.SIGABRT,
		syscall.SIGALRM,
		syscall.SIGBUS,
		syscall.SIGCONT,
		syscall.SIGFPE,
		syscall.SIGHUP,
		syscall.SIGILL,
		syscall.SIGINT,
		syscall.SIGIO,
		syscall.SIGIOT,
		syscall.SIGKILL,
		syscall.SIGPIPE,
		syscall.SIGPROF,
		syscall.SIGQUIT,
		syscall.SIGSEGV,
		syscall.SIGSTOP,
		syscall.SIGSYS,
		syscall.SIGTERM,
		syscall.SIGTRAP,
		syscall.SIGTSTP,
		syscall.SIGTTIN,
		syscall.SIGTTOU,
		syscall.SIGURG,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		syscall.SIGVTALRM,
		syscall.SIGWINCH,
		syscall.SIGXCPU,
		syscall.SIGXFSZ)

	go func() {
		for sig := range s {
			log.Debug(fmt.Sprintf("Signal reveived: %s", sig))

			switch sig {
			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				globalExitCode = 2

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				globalExitCode = 3

			// kill -SIGKILL XXXX
			case syscall.SIGKILL:
				globalExitCode = 9

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				globalExitCode = 15

			default:
				globalExitCode = 1
			}

			ProgramEnd()

		}
	}()
}

func ProgramExit(exitCode int) {
	globalExitCode = 1
	ProgramEnd()
}

func ProgramEnd() {
	log.Info("### program end ###")
	log.Info(fmt.Sprintf("### with exit code %d  ###", globalExitCode))
	os.Exit(globalExitCode)
}
