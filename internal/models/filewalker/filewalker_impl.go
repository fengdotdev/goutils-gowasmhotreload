package filewalker

import (
	"os"
	"path/filepath"
	"time"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/typedef"
)

var _ FileWalkerInterface = &FileWalker{}

// Walk walks the file system and triggers the TriggerFn function if a file is changed
func (fw *FileWalker) Walk(TriggerFn func(path string)) error {

	err := filepath.Walk(fw.root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fw.FileIsIgnored(path) {
			return nil
		}

		if info.IsDir() {
			return nil

		}
		fd, err := typedef.NewFileStatDataFrom(path, info)
		if err != nil {
			return err
		}

		if !fw.CheckFile(path) {
			fw.UpdateFiles(path, *fd)
		}

		fw.DoIfTimeChange(path, fd.LatestModification, func() {
			TriggerFn(path)
		})

		return nil

	})

	return err

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

func (fw *FileWalker) AddIgnoreFile(path string) {
	fw.ignoreFiles[path] = struct{}{}
}

func (fw *FileWalker) AddIgnoreFiles(paths []string) {
	for _, path := range paths {
		fw.ignoreFiles[path] = struct{}{}
	}
}

func (fw *FileWalker) FileIsIgnored(path string) bool {
	if _, ok := fw.ignoreFiles[path]; !ok {
		return false
	}
	return true
}

func (fw *FileWalker) GetRoot() string {
	return fw.root
}

func (fw *FileWalker) GetLastModification(path string) (time.Time, error) {

	if fw.CheckFile(path) {
		return fw.files[path].LatestModification, nil
	}
	return time.Time{}, ErrFileNotFound

}
func (fw *FileWalker) IsSameTime(path string, newtime time.Time) bool {

	last, err := fw.GetLastModification(path)
	if err != nil {
		return false
	}
	return last == newtime
}

func (fw *FileWalker) DoIfTimeChange(path string, newtime time.Time, fn func())  {
	if !fw.IsSameTime(path, newtime) {
		fn()
	}

}
