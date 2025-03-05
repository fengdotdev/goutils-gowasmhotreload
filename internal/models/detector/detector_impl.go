package detector

import (
	"time"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/funcs"
)

var _ DetectorInterface = &Detector{}

func (d *Detector) Detect() error {

	d.DoAndRecover(func() error {
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
	})
	return nil
}

func (d *Detector) Sleep() {
	time.Sleep(d.Settings.GetInterval())
}

func (d *Detector) DoAndRecover(fn func() error) error {
	defer func() {
		if r := recover(); r != nil {
			// log the message: "Panic recovered: <panic message>"
		}
	}()
	fn()
	return nil
}
