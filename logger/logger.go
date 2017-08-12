package logger

import (
	"github.com/op/go-logging"
	"os"
)

// InitLogger initializes the logger
func InitLogger() {
	// set logging format
	var logFormat = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{id:03x} %{module:15s} %{shortfunc:20s} ▶ %{level:.4s} %{color:reset} %{message}`,
	)

	// setup logging backend
	var logBackendInfo = logging.NewLogBackend(os.Stdout, "", 0)
	var logBackendDebug = logging.NewLogBackend(os.Stdout, "", 0)
	var logBackendFormatterInfo = logging.NewBackendFormatter(logBackendInfo, logFormat)
	//var logBackendFormatterDebug = logging.NewBackendFormatter(logBackendDebug, logFormatDebug)
	var logBackendLeveledInfo = logging.AddModuleLevel(logBackendInfo)
	var logBackendLeveledDebug = logging.AddModuleLevel(logBackendDebug)
	logBackendLeveledInfo.SetLevel(logging.INFO, "/")
	logBackendLeveledDebug.SetLevel(-1, "/")

	// Set the backends to be used.
	logging.SetLevel(logging.INFO, "/")
	// logging.SetBackend(logBackendLeveledInfo, logBackendFormatterInfo)
	logging.SetBackend(logBackendLeveledDebug, logBackendFormatterInfo)
}
