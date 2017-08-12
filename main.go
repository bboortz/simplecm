package main

import (
	"github.com/bboortz/simplecm/helper"
	"github.com/bboortz/simplecm/task"
	"github.com/bboortz/simplecm/user"
)

func main() {
	helper.ProgramStart()
	defer helper.ProgramEnd()

	var t = task.NewTask().Build()

	t.OsInstallPackages(
		"a2ps",
		"bash",
		"cups",
		"dosbox",
		"fail2ban",
		"firefox",
		"firefox-beta",
		"git",
		"gcc",
		"gnutls",
		"go",
		"gtk3-print-backends",
		"keepass",
		"libgcrypt",
		"libpulse",
		"libreoffice-fresh",
		"libreoffice-fresh-en-GB",
		"libreoffice-fresh-de",
		"lib32-gnutls",
		"libxml2",
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
		"traceroute",
		"tree",
		"vim",
		"wine",
		"winetricks",
		"wine-mono")
	t.OsUninstallPackages(
		"vi")
	t.OsUpdate()
	t.OsCleanup()
	t.LinkFile("/usr/bin/vim", "/usr/local/bin/vi")

	t.TimeSync()

	t.UserManage("anna", "anna", []string{"anna"})
	t.UserManage("benni", "benni", []string{"benni", "lp", "docker"})
	t.UserManage("userb", "userb", []string{"userb", "lp", "docker", "wheel"})

	user.NewUser().FromUser("userb").Build().BecomeUser()
	t.ShowUser()
	t.HomeManageDirectory("userb")

	user.NewUser().FromUser("benni").Build().BecomeUser()
	t.ShowUser()
	t.HomeManageDirectory("benni")

}
