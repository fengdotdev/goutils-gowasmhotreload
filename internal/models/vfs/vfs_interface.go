package vfs






type VirtualFileSystemInterface interface {
	Init() error
	Walk(TriggerFn func(path string), ignoreFilePaths []string) error
}
