package settings

import "time"

type SettingsInterface interface {
	GetDir() string
	GetInterval() time.Duration
	GetWasmFile() string
	GetCommand() string
	GetPort() string
	GetLang() string
}
