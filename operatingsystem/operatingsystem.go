package operatingsystem

import (
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("/operatingsystem")

func init() {
	logging.SetLevel(logging.DEBUG, "/operatingsystem")

	var logBackend = logging.NewLogBackend(os.Stdout, "", 0)
	var leveledBackend = logging.AddModuleLevel(logBackend)
	leveledBackend.SetLevel(logging.DEBUG, "")
	logging.SetBackend(logBackend)
}

// OperatingSystem is the main interface
type OperatingSystem interface {
	LogOS()
}

// Builder is the interface for the builder
type Builder interface {
	Build() OperatingSystem
}

type operatingsystemBuilder struct {
	//	osname    string
	//	osversion string
}

// NewOperatingSystem creates a new builder instance
func NewOperatingSystem() Builder {
	return &operatingsystemBuilder{}
}

// Build builds a new OperatingSystem instance
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

// LogOS logs out the operating system
func (o *operatingsystem) LogOS() {
	log.Debugf("%s %s", o.osname, o.osversion)
}
