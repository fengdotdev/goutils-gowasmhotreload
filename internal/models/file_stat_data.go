package models

import "time"

type FileStatData struct {
	LatestModification time.Time
	LastCheckSum       string
	LastSize           int64
	UpdateAt           time.Time
}