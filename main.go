package main

import (
	"./helper"
	"./task"
	"./user"
)

func main() {
	helper.ProgramStart()
	defer helper.ProgramEnd()

	var t = task.NewTask().Build()

	t.OsInstallPackages(
		"bash",
		"dosbox",
		"fail2ban",
		"git",
		"gcc",
		"gnutls",
		"go",
		"libpulse",
		"libreoffice-fresh",
		"libreoffice-fresh-en-GB",
		"libreoffice-fresh-de",
		"lib32-gnutls",
		"lsof",
		"make",
		"mono",
		"ntp",
		"python",
		"python-virtualenv",
		"python-pip",
		"q4wine",
		"screen",
		"strace",
		"tree",
		"vim",
		"wine",
		"winetricks",
		"wine-mono")
	t.OsUninstallPackages(
		"vi")
	t.OsUpdate()
	t.OsCleanup()

	t.TimeSync()

	t.UserManage("anna", "anna", []string{"anna"})
	t.UserManage("benni", "benni", []string{"benni"})
	t.UserManage("userb", "userb", []string{"userb", "wheel"})

	user.NewUser().FromUser("userb").Build().BecomeUser()
	t.ShowUser()
	t.HomeManageDirectory("userb")

	user.NewUser().FromUser("benni").Build().BecomeUser()
	t.ShowUser()
	t.HomeManageDirectory("benni")

}
