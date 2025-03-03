package vfs

import "github.com/fengdotdev/goutils-gowasmhotreload/internal/settings"

func NewVirtualFileSystem(root string) *VirtualFileSystem {
	return &VirtualFileSystem{
		root:  root,
		Files: make(map[string]FileStatData),
	}
}

func NewVirtualFileSystemFromSettings(settings *settings.SettingsTemplateInterface) *VirtualFileSystem {
	return NewVirtualFileSystem((*settings).GetDir())
}
