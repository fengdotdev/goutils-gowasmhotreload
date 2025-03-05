package settings

import "time"



type Settings struct {
	Lang     string
	Interval time.Duration
	Dir      string
	WasmFile string
	Command  string
	Port     string
}
