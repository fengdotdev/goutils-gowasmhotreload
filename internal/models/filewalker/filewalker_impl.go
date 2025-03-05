package filewalker

import (

)





var _ FileWalkerInterface = &FileWalker{}


// Walk walks the file system and triggers the TriggerFn function if a file is changed
func (fw *FileWalker) Walk(TriggerFn func(path string), ignoreFilePaths []string) error {
	return nil
}

// UpdateFiles updates the file data in the list
func (fw *FileWalker) UpdateFiles(path string, filedata FileStatData) error {
	if _, ok := fw.files[path]; !ok {
		return ErrFileNotFound
	}
	fw.files[path] = filedata
	return nil
}
// CheckFile checks if the file exists in the list
func (fw *FileWalker) CheckFile(path string) bool {
	if _, ok := fw.files[path]; !ok {
		return false
	}
	return true
}

