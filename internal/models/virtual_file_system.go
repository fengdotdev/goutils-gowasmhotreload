package models

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/funcs"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/msg"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/settings"
)

type VirtualFileSystem struct {
	root string
	// key is the file path
	Files map[string]FileStatData
}

func NewVirtualFileSystem() *VirtualFileSystem {
	return &VirtualFileSystem{
		root:  settings.Settings.Dir,
		Files: make(map[string]FileStatData),
	}
}

func (vfs *VirtualFileSystem) Init() error {
	err := filepath.Walk(vfs.root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			checksum, err := funcs.CalculateChecksum(path)
			if err != nil {
				return err
			}

			vfs.Files[path] = FileStatData{
				LatestModification: info.ModTime(),
				LastSize:           info.Size(),
				LastCheckSum:       checksum,
				UpdateAt:           time.Now(),
			}
		}

		return nil
	})

	return err
}

func (vfs *VirtualFileSystem) Walk(TriggerFn func(path string), ignoreFilePaths []string) error {
	ignoreFiles := make(map[string]string)
	for _, ignoredpath := range ignoreFilePaths {
		ignoreFiles[ignoredpath] = ignoredpath
	}
	err := filepath.Walk(vfs.root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		 // update the file if it is not in the list
		if _, ok := vfs.Files[path]; !ok {

			checksum, err := funcs.CalculateChecksum(path)
			if err != nil {
				return err
			}

			vfs.Files[path] = FileStatData{
				LatestModification: info.ModTime(),
				LastSize:           info.Size(),
				LastCheckSum:       checksum,
				UpdateAt:           time.Now(),
			}
		}

		 // execute the trigger function if the file is not in the ignored list

		somethingChange := false
		whatChanged := ""
		if !info.IsDir() {

			checksum, err := funcs.CalculateChecksum(path)
			if err != nil {
				return err
			}
			if vfs.Files[path].LastCheckSum != checksum {

				somethingChange = true
				whatChanged = path

			}

			if somethingChange {
				vfs.Files[path] = FileStatData{
					LatestModification: info.ModTime(),
					LastSize:           info.Size(),
					LastCheckSum:       checksum,
					UpdateAt:           time.Now(),
				}
				fmt.Println(msg.MSG.ChangeDetectedIn, path, whatChanged)
				TriggerFn(path)
			}
		}

		return nil
	})

	return err
}
