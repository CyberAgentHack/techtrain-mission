package gameapi

import (
	"log"
	"os"
)

// Logger is interface for two levels logging
type Logger interface {
	Warnf(string, ...interface{})
	Debugf(string, ...interface{})
}

// NewLogger returns the pointer for logger
func NewLogger() *Logger {
	return &logger
}

// RawLogger is used to store raw *log.Logger
type RawLogger struct {
	Logger *log.Logger
}

// Logger exports the pointer log.Logger
// If you use other logger, replace logger in `func init()`
var logger Logger = &RawLogger{
	Logger: log.New(os.Stderr, "", log.LstdFlags),
}

// SetLogger sets given logger
// If you use other logger, replace with this function
func SetLogger(l Logger) {
	logger = l
}

// Warnf is a serious log for warning
func (rl *RawLogger) Warnf(format string, v ...interface{}) {
	rl.Logger.Printf(format, v...)
}

// Debugf is a log for debug
// it is not output to *log.Logger
func (rl *RawLogger) Debugf(format string, v ...interface{}) {
	// suppress debug logs
}
