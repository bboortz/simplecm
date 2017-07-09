package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

/*
func execCommand(app string, arg ...string) {
	log.Info("CMD: " + app + " " + strings.Join(arg[:], " "))
	cmd := exec.Command(app, strings.Join(arg[:], " "))
	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Error(err.Error())
		fmt.Printf("%s\n", string(stdoutStderr))
		log.Error("exit program")
		programExit(1)
	}

	fmt.Printf("%s\n", string(stdoutStderr))
}
*/


func execCommandWithResult(command string) (int, string, string) {
	var exitCode int = 0
	//var waitStatus syscall.WaitStatus

	// run command
	log.Debug("CMD: " + command)
	cmd := exec.Command("/bin/sh", "-c", command)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	err := cmd.Run()
	if err != nil {
	    log.Fatal(err)
	}

	// convert buffer to string
	stdoutStr := stdoutBuf.String()
	stderrStr := stderrBuf.String()

	// retrieve exit code
	waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode = waitStatus.ExitStatus()


	return exitCode, stdoutStr, stderrStr
}

func execCommandWithOutput(command string) {
	// run command
	log.Debug("CMD: " + command)
	cmd := exec.Command("/bin/sh", "-c", command)
	var waitStatus syscall.WaitStatus
	var exitCode int
	stdoutStderr, err := cmd.CombinedOutput()

	// retrieve exit code
	waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode = waitStatus.ExitStatus()
	log.Debug(fmt.Sprintf("EXIT CODE: %d", exitCode))

	// print error
	if err != nil {
		fmt.Printf("%s\n", string(stdoutStderr))
		log.Error("ERROR: " + err.Error())
		log.Error(fmt.Sprintf("program exited with exit code %d", exitCode))
		programExit(exitCode)
	}

	fmt.Printf("%s\n", string(stdoutStderr))
}

func execCommandWithExitCode(command string) int {
	var result int = 0
	// run command
	log.Debug("CMD: " + command)
	cmd := exec.Command("/bin/sh", "-c", command)
	var waitStatus syscall.WaitStatus
	cmd.Run()

	// retrieve exit code
	waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
	result = waitStatus.ExitStatus()
	log.Debug(fmt.Sprintf("EXIT CODE: %d", result))

	return result
}
