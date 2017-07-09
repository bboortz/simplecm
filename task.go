package main

import (
	"strings"
)


func osUpdate() {
	log.Info("TASK: " + getFuncName())
	switch detectedOS {
		case "linux-debian":
			execCommandWithOutput("apt-get update")
			execCommandWithOutput("unattended-upgrades -v")
			execCommandWithOutput("apt-get -y dist-upgrade")
		case "linux-archlinux":
			execCommandWithOutput("pacman --noconfirm -Syu")
		default:
			log.Error("Unknown OS detected: " + detectedOS)
			programExit(1)
	}
}

func osInstallPackages(packages ...string) {
	log.Info("TASK: " + getFuncName())
	switch detectedOS {
		case "linux-debian":
			execCommandWithOutput("apt-get update")
			execCommandWithOutput("apt-get -y install " + strings.Join(packages[:], " "))
		case "linux-archlinux":
			execCommandWithOutput("pacman --noconfirm -Sy")
			execCommandWithOutput("pacman --noconfirm -S --needed " + strings.Join(packages[:], " "))
		default:
			log.Error("Unknown OS detected: " + detectedOS)
			programExit(1)
	}
}

func osCleanup(packages ...string) {
	log.Info("TASK: " + getFuncName())
	switch detectedOS {
		case "linux-debian":
			execCommandWithOutput("apt-get -y autoremove")
			execCommandWithOutput("apt-get -y autoclean")
			execCommandWithOutput("apt-get -y clean")
		case "linux-archlinux":
			execCommandWithOutput("rm -f /var/cache/pacman/pkg/*")
//			execCommandWithOutput("pacman -Rns $(pacman -Qtdq)")
		default:
			log.Error("Unknown OS detected: " + detectedOS)
			programExit(1)
	}
}

func timeSync() {
	log.Info("TASK: " + getFuncName())
	switch detectedOS {
		case "linux-archlinux":
			execCommandWithOutput("ln -sfn /usr/share/zoneinfo/Europe/Berlin /etc/localtime")
			execCommandWithOutput("systemctl enable ntpd")
			execCommandWithOutput("systemctl start ntpd")
		default:
			log.Error("Unknown OS detected: " + detectedOS)
			programExit(1)
	}
}

func userManage(username string, loginGroup string, groups []string) {
	log.Info("TASK: " + getFuncName())
	groupManage(loginGroup)
	for _, g := range groups {
		groupManage(g)
	}
	userPresent := helperUserPresent(username)
	if !userPresent {
		execCommandWithOutput("useradd --create-home --gid " + loginGroup + " " + username)
	}
	var groupCommaSep string = strings.Join(groups[:], ",")
	execCommandWithOutput("usermod --gid " + loginGroup + " --groups " + groupCommaSep + " " + loginGroup)
}

func groupManage(groupname string) {
	log.Info("TASK: " + getFuncName())
	groupPresent := helperGroupPresent(groupname)
	if !groupPresent {
		execCommandWithOutput("groupadd " + groupname)
	}
}


func homeManageDirectory(username string) {
	log.Info("TASK: " + getFuncName())
	execCommandWithOutput("chmod 700 ~" + username)
	execCommandWithOutput("mkdir -p ~" + username + "/.ssh && chmod 700 ~" + username + "/.ssh")
	execCommandWithOutput("mkdir -p ~" + username + "/bin")
	execCommandWithOutput("mkdir -p ~" + username + "/files")
	execCommandWithOutput("mkdir -p ~" + username + "/progs")
}

