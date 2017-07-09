package main

import (
)


var debug bool = false
var globalExitCode = 0
var detectedOS string


func main() {
	programStart()
	defer programEnd()

	userCheckRoot()

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
		"tree",
//		unattended-upgrades",
		"vim")
	osUpdate()
	osCleanup()

	timeSync()

	userManage("anna","anna", []string{"anna"})
	userManage("benni","benni", []string{"benni"})
	userManage("userb","userb", []string{"userb", "wheel"})

	setUserPrivileges("benni", "benni", "")
	homeManageDirectory("benni")
	setUserPrivileges("userb", "userb", "")
	homeManageDirectory("userb")
	userCheckRoot()
}
