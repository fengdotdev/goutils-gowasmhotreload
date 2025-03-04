package detector

import (
	"time"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/funcs"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/models/vfs"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/settings"
)

func NewDetector(s settings.SettingsTemplateInterface, v vfs.VirtualFileSystemInterface) *Detector {
	return &Detector{
		Settings:        s,
		Vfs:             v,
		ignoredFiles:    make([]string, 0),
		counter:         0,
		FileThatChanged: "",
		somethingChange: false,
	}
}

type Detector struct {
	Settings        settings.SettingsTemplateInterface
	Vfs             vfs.VirtualFileSystemInterface
	ignoredFiles    []string
	counter         int
	FileThatChanged string
	somethingChange bool
}

func (d *Detector) Detect() error {

	defer func() error {
		if r := recover(); r != nil {
			// log the message: "Error: <error>"
			return r.(error)
		}
		return nil
	}()

	for {
		time.Sleep(d.Settings.GetInterval())
		if !d.somethingChange && d.counter == 0 {

			// log the message: "Watching for changes in <lang> <dir>"
		}
		d.counter++
		// Check for changes in the directory
		d.Vfs.Walk(func(path string) {
			d.somethingChange = true
			d.FileThatChanged = path
		}, d.ignoredFiles)

		if d.somethingChange {
			//

			err := funcs.CommandExec(d.Settings.GetCommand())
			if err != nil {

			} else {

			}
			// Reset the state
			d.somethingChange = false
			d.FileThatChanged = ""
			d.counter = 0
		}
	}
}
