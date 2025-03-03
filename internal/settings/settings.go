package settings

import "time"

type SettingsTemplateInterface interface {
	GetDir() string
	GetInterval() time.Duration
	GetWasmFile() string
	GetCommand() string
	GetPort() string
}

var _ SettingsTemplateInterface = &SettingsTemplate{}

type SettingsTemplate struct {
	Interval time.Duration
	Dir      string
	WasmFile string
	Command  string
	Port     string
}
