package settings

import "time"

var _ SettingsInterface = &Settings{}

func (s *Settings) GetDir() string {
	return s.Dir
}

func (s *Settings) GetInterval() time.Duration {
	return s.Interval
}

func (s *Settings) GetWasmFile() string {
	return s.WasmFile
}

func (s *Settings) GetCommand() string {
	return s.Command
}

func (s *Settings) GetPort() string {
	return s.Port
}

func (s *Settings) GetLang() string {
	return s.Lang
}
