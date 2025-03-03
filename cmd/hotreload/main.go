package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/funcs"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/models/vfs"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/msg"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/server"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/settings"
)

func main() {

	df := settings.DefaultSettings

	log.Println(msg.MSG.Welcome, df.Dir)
	ignoredFiles := []string{df.WasmFile}
	somethingChange := false
	FilenameThatChanged := ""
	counter := 0
	vfs := vfs.NewVirtualFileSystemFromSettings(df)
	err := vfs.Init()
	if err != nil {
		panic(err)
	}

	svr := server.NewServer()

	go svr.Start()
	for {
		time.Sleep(settings.Settings.Interval)
		if !somethingChange && counter == 0 {
			fmt.Println(msg.MSG.WatchingChanges, settings.Settings.Dir)
		}
		counter++
		// Check for changes in the directory
		vfs.Walk(func(path string) {
			somethingChange = true
			FilenameThatChanged = path
		}, ignoredFiles)

		if somethingChange {
			fmt.Println(msg.MSG.ChangeDetectedIn, FilenameThatChanged)

			err := funcs.CommandExec(settings.Settings.Command)
			if err != nil {
				fmt.Println(msg.MSG.MsgError, err)
			} else {
				fmt.Println(msg.MSG.ExecCommand, settings.Settings.Command)
				fmt.Println(msg.MSG.SusscessfulCompilation)
			}
			// Reset the state
			somethingChange = false
			FilenameThatChanged = ""
			counter = 0
		}
	}
}
