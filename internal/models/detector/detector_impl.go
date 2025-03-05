package detector

import (
	"time"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/funcs"
)

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
