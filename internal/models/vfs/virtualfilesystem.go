package vfs




type VirtualFileSystem struct {
	root string
	// key is the file path
	Files map[string]FileStatData
}

