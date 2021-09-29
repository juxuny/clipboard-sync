package log

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewPrefix("test")
	logger.Info("this is test log.")
	Info("test log")
	Debug("debug")
}
