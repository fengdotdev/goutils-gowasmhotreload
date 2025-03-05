package typedef

import (
	"os"
	"time"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/funcs"
)

type FileStatData struct {
	LatestModification time.Time
	LastCheckSum       string
	LastSize           int64
	UpdateAt           time.Time
}


// NewFileStatDataFrom creates a new FileStatData from the given path and os.FileInfo or returns an error
func NewFileStatDataFrom(path string, info os.FileInfo) (*FileStatData, error) {

	checksum, err := funcs.CalculateChecksum(path)
	if err != nil {
		return nil, err
	}

	return &FileStatData{
		LatestModification: info.ModTime(),
		LastCheckSum:       checksum,
		LastSize:           info.Size(),
		UpdateAt:           time.Now(),
	}, nil
}
