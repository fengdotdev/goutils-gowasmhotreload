package filewalker



type FileWalkerInterface interface {
	Walk(TriggerFn func(path string), ignoreFilePaths []string) error
}