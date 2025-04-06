package executor

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

// GetApps func
func GetApps() *[]App {
	fmt.Println("imported Apps:", apps)
	return &apps
}

// ImportApp func
func ImportApp(app App) {
	err := os.Chdir(rpdPath)
	if err != nil {
		fmt.Println("Error changing to RPD Path:", err)
		return
	}

	f, err := os.Create("importedApps/" + app.Name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer f.Close()

	if app.ProtonPath == "" {
		app.ProtonPath = "none"
	}

	if app.WinePrefixPath == "" {
		app.WinePrefixPath = "none"
	}

	if app.CompatDataPath == "" {
		app.CompatDataPath = "none"
	}

	f.WriteString(app.Name + "\n")
	f.WriteString(app.GamePath + "\n")
	f.WriteString(app.ProtonPath + "\n")
	f.WriteString(app.WinePrefixPath + "\n")
	f.WriteString(app.CompatDataPath + "\n")
	if app.Args != "" {
		f.WriteString(app.Args)
	}

	f.Sync()

	apps = append(apps, app)
	sortApps()
}

// EditApp func
func EditApp(oldApp, newApp App) {
	DeleteApp(oldApp)
	ImportApp(newApp)
}

// DeleteApp func
func DeleteApp(app App) {
	err := os.Chdir(rpdPath)
	if err != nil {
		fmt.Println("Error changing to RPD Path:", err)
		return
	}

	err = os.Remove("importedApps/" + app.Name)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(apps); i++ {
		if apps[i] == app {
			apps = append(apps[:i], apps[i+1:]...)
		}
	}
}

func loadImportedApps() []App {
	files, err := ioutil.ReadDir("importedApps")
	if err != nil {
		log.Fatal(err)
	}

	loadedApps := make([]App, 0)

	for _, fileMeta := range files {
		file, err := os.Open("importedApps/" + fileMeta.Name())
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(file)

		config := make([]string, 0)

		for scanner.Scan() {
			config = append(config, scanner.Text())
		}
		cleanStrings(&config)

		newApp := App{Name: config[0], GamePath: config[1], ProtonPath: config[2], WinePrefixPath: config[3], CompatDataPath: config[4]}
		if len(config) == 6 {
			newApp.Args = config[5]
		}

		loadedApps = append(loadedApps, newApp)

		file.Close()
	}
	apps = loadedApps
	return loadedApps
}

func sortApps() {
	sort.Slice(apps, func(i, j int) bool {
		return apps[i].Name < apps[j].Name
	})
}

func checkForDataFolder() bool {
	if _, err := os.Stat("importedApps"); !os.IsNotExist(err) {
		return true
	}
	os.Mkdir("importedApps", os.ModePerm)
	return false
}

func cleanStrings(str *[]string) {
	for i := 0; i < len(*str); i++ {
		(*str)[i] = strings.Replace((*str)[i], "\n", "", -1)
		(*str)[i] = strings.Replace((*str)[i], "none", "", -1)
	}
}
