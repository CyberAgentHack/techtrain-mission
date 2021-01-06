package log

import (
	"fmt"
	"log"
	"os"

	"github.com/task4233/techtrain-mission/gameapi/config"
)

// Logger is interface for two levels logging
type Logger interface {
	Warnf(string, ...interface{})
	Debugf(string, ...interface{})
}

// RawLogger is used to store raw *log.Logger
type RawLogger struct {
	Logger *log.Logger
}

// SetDebugStatus sets whether debug is on or not
func SetDebugStatus(status bool) {
	debugOn = status
}

var debugOn = config.IsDev()

// MyLogger exports the pointer log.Logger
// If you use other logger, replace logger in `func init()`
var MyLogger Logger = &RawLogger{
	Logger: log.New(os.Stderr, "", log.LstdFlags),
}

// SetLogger sets given logger
// If you use other logger, replace with this function
func SetLogger(l Logger) {
	MyLogger = l
}

// Warnf is a serious log for warning
func (rl *RawLogger) Warnf(format string, v ...interface{}) {
	rl.Logger.Printf(format, v...)
}

// Debugf is a log for debug
// it is not output to *log.Logger
func (rl *RawLogger) Debugf(format string, v ...interface{}) {
	if debugOn {
		rl.Logger.Printf(fmt.Sprintf("[DEBUG]: %s", format), v...)
	}
}
