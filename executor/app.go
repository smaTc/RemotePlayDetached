package executor

// App struct
type App struct {
	Name           string
	GamePath       string
	Args           string
	ProtonPath     string
	WinePrefixPath string
	CompatDataPath string
}

// Equals func
func (a *App) Equals(otherApp App) bool {
	return a.Name == otherApp.Name &&
		a.GamePath == otherApp.GamePath &&
		a.Args == otherApp.Args &&
		a.ProtonPath == otherApp.ProtonPath &&
		a.WinePrefixPath == otherApp.WinePrefixPath &&
		a.CompatDataPath == otherApp.CompatDataPath
}
