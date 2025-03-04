package settings

import "time"

type SettingsTemplateInterface interface {
	GetDir() string
	GetInterval() time.Duration
	GetWasmFile() string
	GetCommand() string
	GetPort() string
	GetLang() string
}


type SettingsTemplate struct {
	Lang     string
	Interval time.Duration
	Dir      string
	WasmFile string
	Command  string
	Port     string
}
