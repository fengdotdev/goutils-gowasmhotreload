package filewalker

type FileWalker struct {
	root string
	// key is the file path
	files       map[string]FileStatData
	ignoreFiles map[string]string
}
