package operatingsystem

import (
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("/operatingsystem")

func init() {
	logging.SetLevel(logging.INFO, "/operatingsystem")

	var logBackend = logging.NewLogBackend(os.Stdout, "", 0)
	var leveledBackend = logging.AddModuleLevel(logBackend)
	leveledBackend.SetLevel(logging.NOTICE, "")
	logging.SetBackend(logBackend)
}

type OperatingSystem interface {
	LogOS()
}

type OperatingSystemBuilder interface {
	Build() OperatingSystem
}

type operatingsystemBuilder struct {
	osname    string
	osversion string
}

func NewOperatingSystem() OperatingSystemBuilder {
	return &operatingsystemBuilder{}
}

func (b *operatingsystemBuilder) Build() OperatingSystem {
	result := &operatingsystem{
		osname:    detectOS(),
		osversion: "unknown",
	}

	result.LogOS()
	return result
}

type operatingsystem struct {
	osname    string
	osversion string
}

func (o *operatingsystem) LogOS() {
	log.Debugf("%s %s", o.osname, o.osversion)
}
