package main

import (
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/models/detector"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/models/filewalker"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/server"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/settings"
)

func main() {

	// grab flags

	// setup

	defaultSettings := settings.NewSettingsFromDefaults()

	// start server

	devserver := server.NewServerFromSettings(&defaultSettings)
	devserver.StartNoBlock()

	// start walker

	walker := filewalker.NewFileWalker(defaultSettings.Dir)

	// start detector

	detec := detector.NewDetector(&defaultSettings, walker)

	detec.Detect()

}
