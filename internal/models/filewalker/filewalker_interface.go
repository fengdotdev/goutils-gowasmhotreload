package filewalker

type FileWalkerInterface interface {
	Walk(TriggerFn func(path string)) error
	UpdateFiles(path string, filedata FileStatData) error
	CheckFile(path string) bool
	AddIgnoreFile(path string)
	AddIgnoreFiles(paths []string)
	FileIsIgnored(path string) bool
	GetRoot() string
}
