package main

import (
)


var debug bool = false
var globalExitCode = 0
var detectedOS string


func main() {
	programStart()
	defer programEnd()

	osInstallPackages(
		"bash",
		"fail2ban",
		"git",
		"gcc",
		"libreoffice-fresh",
		"libreoffice-fresh-en-GB",
		"libreoffice-fresh-de",
		"lsof",
		"make",
		"ntp",
		"python",
		"python-virtualenv",
		"python-pip",
		"screen",
		"strace",
//		unattended-upgrades",
		"vim")
	osUpdate()
	osCleanup()

	timeSync()

	userManage("anna")
	userManage("benni")
	userManage("testuser")
	homeManageDirectory("benni")
}
