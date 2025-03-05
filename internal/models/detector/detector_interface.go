package detector

type DetectorInterface interface {
	Detect() error
	Sleep()
	DoAndRecover(fn func() error) error
}
