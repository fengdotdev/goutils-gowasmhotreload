package settings

import "time"

var DefaultSettings = SettingsTemplate{
	Interval: 1 * time.Second,
	Dir:      "./",
	WasmFile: "main.wasm",
	Command:  "GOOS=js GOARCH=wasm go build -o main.wasm",
	Port:     ":5656",
}
