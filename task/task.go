package task

import (
	"../helper"
	"../operatingsystem"
	"../reflect"
	"github.com/op/go-logging"
	"strings"
)

var log = logging.MustGetLogger("task")
var detectedOS string = "linux-archlinux"

func init() {
	logging.SetLevel(logging.INFO, "")
}

type Task interface {
	OsUpdate()
	OsInstallPackages(packages ...string)
	OsCleanup(packages ...string)
	TimeSync()
	ShowUser()
	UserManage(username string, loginGroup string, groups []string)
	HomeManageDirectory(username string)
}

type TaskBuilder interface {
	Build() Task
}

type taskBuilder struct {
	os operatingsystem.OperatingSystem
}

func NewTask() TaskBuilder {
	return &taskBuilder{}
}

func (b *taskBuilder) Build() Task {
	return &task{
		os: operatingsystem.NewOperatingSystem().Build(),
	}
}

type task struct {
	os operatingsystem.OperatingSystem
}

func (o *task) OsUpdate() {
	log.Info("TASK: " + reflect.GetFuncName())
	switch detectedOS {
	case "linux-debian":
		operatingsystem.ExecCommand("apt-get update")
		operatingsystem.ExecCommand("unattended-upgrades -v")
		operatingsystem.ExecCommand("apt-get -y dist-upgrade")
	case "linux-archlinux":
		operatingsystem.ExecCommand("pacman --noconfirm -Syu")
	default:
		log.Error("Unknown OS detected: " + detectedOS)
		helper.ProgramExit(1)
	}
}

func (o *task) OsInstallPackages(packages ...string) {
	log.Info("TASK: " + reflect.GetFuncName())
	switch detectedOS {
	case "linux-debian":
		operatingsystem.ExecCommand("apt-get update")
		operatingsystem.ExecCommand("apt-get -y install " + strings.Join(packages[:], " "))
	case "linux-archlinux":
		operatingsystem.ExecCommand("pacman --noconfirm -Sy")
		operatingsystem.ExecCommand("pacman --noconfirm -S --needed " + strings.Join(packages[:], " "))
	default:
		log.Error("Unknown OS detected: " + detectedOS)
		helper.ProgramExit(1)
	}
}

func (o *task) OsCleanup(packages ...string) {
	log.Info("TASK: " + reflect.GetFuncName())
	switch detectedOS {
	case "linux-debian":
		operatingsystem.ExecCommand("apt-get -y autoremove")
		operatingsystem.ExecCommand("apt-get -y autoclean")
		operatingsystem.ExecCommand("apt-get -y clean")
	case "linux-archlinux":
		operatingsystem.ExecCommand("rm -f /var/cache/pacman/pkg/*")
		//			operatingsystem.ExecCommand("pacman -Rns $(pacman -Qtdq)")
	default:
		log.Error("Unknown OS detected: " + detectedOS)
		helper.ProgramExit(1)
	}
}

func (o *task) TimeSync() {
	log.Info("TASK: " + reflect.GetFuncName())
	switch detectedOS {
	case "linux-archlinux":
		operatingsystem.ExecCommand("ln -sfn /usr/share/zoneinfo/Europe/Berlin /etc/localtime")
		operatingsystem.ExecCommand("systemctl enable ntpd")
		operatingsystem.ExecCommand("systemctl start ntpd")
	default:
		log.Error("Unknown OS detected: " + detectedOS)
		helper.ProgramExit(1)
	}
}

func (o *task) ShowUser() {
	log.Info("TASK: " + reflect.GetFuncName())
	operatingsystem.ExecCommand("id")
}

func (o *task) UserManage(username string, loginGroup string, groups []string) {
	log.Info("TASK: " + reflect.GetFuncName())
	o.GroupManage(loginGroup)
	for _, g := range groups {
		o.GroupManage(g)
	}
	userPresent := helper.IsUserPresent(username)
	if !userPresent {
		operatingsystem.ExecCommand("useradd --create-home --gid " + loginGroup + " " + username)
	}
	var groupCommaSep string = strings.Join(groups[:], ",")
	operatingsystem.ExecCommand("usermod --gid " + loginGroup + " --groups " + groupCommaSep + " " + loginGroup)
}

func (o *task) GroupManage(groupname string) {
	log.Info("TASK: " + reflect.GetFuncName())
	groupPresent := helper.IsGroupPresent(groupname)
	if !groupPresent {
		operatingsystem.ExecCommand("groupadd " + groupname)
	}
}

func (o *task) HomeManageDirectory(username string) {
	log.Info("TASK: " + reflect.GetFuncName())
	operatingsystem.ExecCommand("chmod 700 ~" + username)
	operatingsystem.ExecCommand("mkdir -p ~" + username + "/.ssh && chmod 700 ~" + username + "/.ssh")
	operatingsystem.ExecCommand("mkdir -p ~" + username + "/bin")
	operatingsystem.ExecCommand("mkdir -p ~" + username + "/files")
	operatingsystem.ExecCommand("mkdir -p ~" + username + "/progs")
}
