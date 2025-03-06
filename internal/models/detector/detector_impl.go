package detector

import (
	"time"
)

var _ DetectorInterface = &Detector{}

func (d *Detector) Detect() error {

	for {
		d.Sleep()
		d.filewalker.Walk()
	}
	return nil
}

func (d *Detector) Sleep() {
	time.Sleep(d.settings.GetInterval())
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
