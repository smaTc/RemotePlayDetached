package fynegui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/smaTc/RemotePlayDetached/executor"
)

var apps *[]executor.App

var appList *widget.Form

//var appList widget.NewVBox
var appListContainer *fyne.Container

func importApp() {
	importWindow := rpd.NewWindow("Import App")
	importWindow.Resize(fyne.NewSize(400, 150))

	name := widget.NewFormItem("Name", widget.NewEntry())
	path := widget.NewFormItem("Path", widget.NewEntry())
	args := widget.NewFormItem("Args", widget.NewEntry())
	form := widget.NewForm(name, path, args)

	cancelButton := widget.NewButton("Cancel", func() {
		importWindow.Close()
	})

	okButton := widget.NewButton("OK", func() {
		appName := name.Widget.(*widget.Entry).Text
		appPath := path.Widget.(*widget.Entry).Text
		argsString := args.Widget.(*widget.Entry).Text

		if appName == "" || appPath == "" {
			return
		}

		newApp := executor.App{Name: appName, Path: appPath, Args: argsString}
		importWindow.Close()
		executor.ImportApp(newApp)
		refreshContent()
	})

	buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), okButton, layout.NewSpacer(), cancelButton)

	importWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), form, layout.NewSpacer(), buttons))

	importWindow.Show()
}

func buildAppListContainer() *fyne.Container {
	appListContainer = loadApps()
	return appListContainer
}

func addToList(app executor.App) {
	appItem := NewAppItem(app)
	formItem := widget.NewFormItem(app.Name, appItem)
	appList.AppendItem(formItem)
}

func loadApps() *fyne.Container {
	counter := 1
	fmt.Println("(re)loading apps")
	apps = executor.GetApps()
	itemList := make([]*widget.FormItem, 0)

	for _, app := range *apps {
		appItem := NewAppItem(app)
		//formItem := widget.NewFormItem(app.Name, appItem)
		formItem := widget.NewFormItem(strconv.Itoa(counter)+". "+app.Name, appItem)
		counter++
		itemList = append(itemList, formItem)
	}

	appList = widget.NewForm(itemList...)
	appGroup := widget.NewGroup("Apps", appList)
	appListContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), appGroup)

	//appScroller := widget.NewScrollContainer(appList)
	//appScroller.Resize(fyne.NewSize(850, 490))
	//appListContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), appScroller)

	return appListContainer
}
