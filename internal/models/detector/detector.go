package detector

type Detector struct {
	Settings        Settings
	Vfs             FileWalker
	ignoredFiles    []string
	counter         int
	FileThatChanged string
	somethingChange bool
}
