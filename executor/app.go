package executor

//App struct
type App struct {
	Name string
	Path string
	Args string
}

//Equals func
func (a *App) Equals(otherApp App) bool {
	return a.Name == otherApp.Name && a.Path == otherApp.Path && a.Args == otherApp.Args
}
