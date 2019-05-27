package astilog

import (
	"log/syslog"
	"os"

	"github.com/rs/xlog"
)

type xlogger struct {
	xlog.Logger
}

func (l *xlogger) Clone() Logger                     { return l }
func (l *xlogger) WithField(k string, v interface{}) {}
func (l *xlogger) WithFields(fs Fields)              {}

// NewXLogger returns a new xlog.Logger
func NewXLogger(c Configuration) Logger {
	return &xlogger{
		Logger: xlog.New(NewXLoggerConfig(c)),
	}
}

// NewXLoggerConfig returns a new xlog.Config
func NewXLoggerConfig(c Configuration) (o xlog.Config) {
	// Build config
	o = xlog.Config{
		Fields: xlog.F{
			"app_name": c.AppName,
		},
		Level:  xlog.LevelInfo,
		Output: xlog.NewConsoleOutputW(os.Stderr, xlog.NewSyslogOutputFacility("", "", c.AppName, syslog.LOG_LOCAL0)),
	}

	// Verbose
	if c.Verbose {
		o.Level = xlog.LevelDebug
	}
	return
}
