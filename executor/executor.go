package executor

import (
	"errors"
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

//RunApp func
func RunApp(app App) error {
	return executeApp(app)
}

//RunAppWithArgs func
func RunAppWithArgs(mode, app string) error {
	switch mode {
	case "direct":
		return executeApp(App{Path: app})
	case "list":
		var found bool = false
		for _, listApp := range apps {
			if listApp.Name == app {
				found = true
				return executeApp(listApp)
			}
		}
		if !found {
			fmt.Println(app, "not found in Appliaction list")
			return errors.New("app not found")
		}
	}
	return nil
}
