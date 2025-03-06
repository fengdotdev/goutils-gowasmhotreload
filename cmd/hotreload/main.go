package main

import (
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/models/detector"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/models/filewalker"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/server"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/settings"
)



// hot reload compila go wasm y recarga el navegador
// 1. detecta cambios en el directorio 
// 2. compila go wasm
// 3. recarga el navegador
// 4. repite


// que necesitamos?
// directorio a observar
// intervalo de tiempo para observar
// comando para compilar go wasm
// comando para recargar el navegador


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
