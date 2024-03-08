package loggerMy

import (
	"testing"
)

func TestLoggerPrint(t *testing.T) {
	Init("traceid", "iptrace")
	Debug("Debug message")
	Debug("Debug message2")
	Info("Info message")
	Warn("Warn message")
	Print()
}

func TestLoggerError(t *testing.T) {
	Init("traceid", "iptrace")
	Debug("Debug message")
	Debug("Debug message2")
	Info("Info message")
	Info("Info message2")
	Error("code", "message", "trace")
}
