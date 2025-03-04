package vfs

type VirtualFileSystem struct {
	lang string
	root string
	// key is the file path
	Files map[string]FileStatData
}
