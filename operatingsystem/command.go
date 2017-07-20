package operatingsystem

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

func ExecCommand(command string) (int, string, string) {

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
	exitCode := waitStatus.ExitStatus()
	log.Debug(fmt.Sprintf("EXIT CODE: %d", exitCode))

	return exitCode, stdoutStr, stderrStr
}
