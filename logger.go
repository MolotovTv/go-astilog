package astilog

import "github.com/rs/xlog"

// NewConfig creates a new xlog.Config
func NewConfig(c Configuration) (o xlog.Config) {
	// Init
	o = xlog.Config{
		Fields: xlog.F{
			"app_name": c.AppName,
		},
		Level:  xlog.LevelInfo,
		Output: DefaultOutput(c),
	}

	// Verbose
	if c.Verbose {
		o.Level = xlog.LevelDebug
	}
	return
}

// New creates a new xlog.Logger
func New(c Configuration) xlog.Logger {
	return xlog.New(NewConfig(c))
}
