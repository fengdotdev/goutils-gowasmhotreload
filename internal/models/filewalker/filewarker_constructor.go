package filewalker

func NewFileWalker(root string) *FileWalker {
	return &FileWalker{
		root:        root,
		files:       make(map[string]FileStatData),
		ignoreFiles: make(map[string]string),
	}
}
