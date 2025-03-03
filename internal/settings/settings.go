package settings

import "time"


type SettingsTemplate struct {
	Interval time.Duration
	Dir      string
	WasmFile string
	Command  string
	Port     string
}
 

var DefaultSettings = SettingsTemplate{
	Interval:1 * time.Second,
	Dir:"./",
	WasmFile:"main.wasm",
	Command:"GOOS=js GOARCH=wasm go build -o main.wasm",
	Port:":5656",
}


var Settings = DefaultSettings

