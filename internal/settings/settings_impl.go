package settings

import "time"

var _ SettingsTemplateInterface = &SettingsTemplate{}

func (s *SettingsTemplate) GetDir() string {
	return s.Dir
}

func (s *SettingsTemplate) GetInterval() time.Duration {
	return s.Interval
}

func (s *SettingsTemplate) GetWasmFile() string {
	return s.WasmFile
}

func (s *SettingsTemplate) GetCommand() string {
	return s.Command
}

func (s *SettingsTemplate) GetPort() string {
	return s.Port
}

func (s *SettingsTemplate) GetLang() string {
	return s.Lang
}
