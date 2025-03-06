package detector

func NewDetector(settings Settings, filewalker FileWalker) *Detector {
	return &Detector{
		settings:   settings,
		filewalker: filewalker,
	}
}
