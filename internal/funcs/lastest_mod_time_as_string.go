package funcs

import (
	"log"
	"os"
)

func LatestModTimeAsString(path string) string {
	filestat, err := os.Stat(path)
	if err != nil {
		log.Println(err)
	}
	return filestat.ModTime().String()
}
