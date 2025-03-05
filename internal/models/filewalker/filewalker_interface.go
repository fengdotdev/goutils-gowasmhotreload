package filewalker

type FileWalkerInterface interface {
	Walk(TriggerFn func(path string), ignoreFilePaths []string) error
	UpdateFiles(path string, filedata FileStatData) error
	CheckFile(path string) bool
}
