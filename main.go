package main

import (
	"github.com/smaTc/RemotePlayDetached/executor"
	"github.com/smaTc/RemotePlayDetached/fynegui"
)

func main() {
	executor.Init()
	fynegui.Run()
}
