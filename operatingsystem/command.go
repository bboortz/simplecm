package operatingsystem

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

// ExecCommand la
func ExecCommand(command string) (int, string, string) {
	return ExecCommandAllParams(command, true)
}

// ExecCommandWithoutErrCheck la
func ExecCommandWithoutErrCheck(command string) (int, string, string) {
	return ExecCommandAllParams(command, false)
}

// ExecCommandAllParams la
func ExecCommandAllParams(command string, checkError bool) (int, string, string) {

	// run command
	log.Debug("CMD: " + command)
	cmd := exec.Command("/bin/sh", "-c", command)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	err := cmd.Run()
	if checkError && err != nil {
		log.Info("baaem")
		log.Fatal(err)
	}

	// convert buffer to string
	stdoutStr := stdoutBuf.String()
	stderrStr := stderrBuf.String()

	// retrieve exit code
	waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode := waitStatus.ExitStatus()
	log.Debug(fmt.Sprintf("EXIT CODE: %d", exitCode))

	return exitCode, stdoutStr, stderrStr
}
