package main

import (
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("main")

// set logging format
var logFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{id:03x} %{shortfunc:20s} ▶ %{level:.4s} %{color:reset} %{message}`,
)

// log logging backend
var logBackendStdout = logging.NewLogBackend(os.Stdout, "", 0)
var logBackendStderr = logging.NewLogBackend(os.Stderr, "", 0)
var logBackendFormatterStdout = logging.NewBackendFormatter(logBackendStdout, logFormat)
var logBackendFormatterStderr = logging.NewBackendFormatter(logBackendStderr, logFormat)
