package detector

func NewDetector(s Settings, v FileWalker) *Detector {
	return &Detector{
		Settings:        s,
		Vfs:             v,
		ignoredFiles:    make([]string, 0),
		counter:         0,
		FileThatChanged: "",
		somethingChange: false,
	}
}
