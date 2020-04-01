package executor

import (
	"fmt"
	"os"
)

var apps = make([]App, 0)

//Init func
func Init() []App {
	newcwd, err := os.Getwd()
	cwd = newcwd
	if err != nil {
		fmt.Println(err)
	}

	load := checkForDataFolder()
	if load {
		return loadImportedApps()
	}
	return nil
}
