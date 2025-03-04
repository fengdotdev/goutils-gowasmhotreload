package main

import (
	"fmt"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/models/detector"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/models/vfs"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/msg"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/server"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/settings"
)

func main() {

	// setup

	df := settings.DefaultSettings

	msg := msg.GetMSG(df.Lang)

	fmt.Println(msg) // fix

	svr := server.NewServerFromSettings(&df)

	svr.StartNoBlock()

	vfs := vfs.NewVirtualFileSystemFromSettings(&df)
	err := vfs.Init()
	if err != nil {
		panic(err)
	}

	det := detector.NewDetector(&df, vfs)
	det.Detect()

}
