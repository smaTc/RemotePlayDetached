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

//RunApp func
func RunApp(app App) {
	executeApp(app)
}

//RunAppWithArgs func
func RunAppWithArgs(mode, app string) {
	switch mode {
	case "direct":
		executeApp(App{Path: app})
	case "list":
		var found bool = false
		for _, listApp := range apps {
			if listApp.Name == app {
				found = true
				executeApp(listApp)
			}
		}
		if !found {
			fmt.Println(app, "not found in Appliaction list")
			os.Exit(1)
		}
	}

}
