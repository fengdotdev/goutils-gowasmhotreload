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

	vfs := vfs.NewVirtualFileSystemFromSettings(&df)
	err := vfs.Init()
	if err != nil {
		panic(err)
	}

	svr := server.NewServerFromSettings(&df)

	go svr.Start()

	watchForChanges(vfs, &df)
}

func watchForChanges(v *vfs.VirtualFileSystem, s settings.SettingsTemplateInterface) {

	ignoredFiles := []string{s.GetWasmFile()}
	somethingChange := false
	FilenameThatChanged := ""
	counter := 0

	for {
		time.Sleep(s.GetInterval())
		if !somethingChange && counter == 0 {
			fmt.Println(msg.MSG.WatchingChanges, s.GetDir())
		}
		counter++
		// Check for changes in the directory
		v.Walk(func(path string) {
			somethingChange = true
			FilenameThatChanged = path
		}, ignoredFiles)

		if somethingChange {
			fmt.Println(msg.MSG.ChangeDetectedIn, FilenameThatChanged)

			err := funcs.CommandExec(s.GetCommand())
			if err != nil {
				fmt.Println(msg.MSG.MsgError, err)
			} else {
				fmt.Println(msg.MSG.ExecCommand, s.GetCommand())
				fmt.Println(msg.MSG.SusscessfulCompilation)
			}
			// Reset the state
			somethingChange = false
			FilenameThatChanged = ""
			counter = 0
		}
	}
}
