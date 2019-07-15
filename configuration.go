package astilog

// Formats
const (
	FormatJSON = "json"
	FormatText = "text"
)

// Outs
const (
	OutFile   = "file"
	OutStdOut = "stdout"
	OutSyslog = "syslog"
)

// Configuration represents the configuration of the logger
type Configuration struct {
	AppName         string `toml:"app_name"`
	DisableColors   bool   `toml:"disable_colors"`
	Filename        string `toml:"filename"`
	FullTimestamp   bool   `toml:"full_timestamp"`
	Format          string `toml:"format"`
	MessageKey      string `toml:"message_key"`
	Out             string `toml:"out"`
	TimestampFormat string `toml:"timestamp_format"`
	Verbose         bool   `toml:"verbose"`
}
