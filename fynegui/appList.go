package fynegui

import (
	"fmt"

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

	nameEntry := NewButtonEntry()
	pathEntry := NewButtonEntry()
	argsEntry := NewButtonEntry()

	/*
		nameEntry := widget.NewEntry()
		pathEntry := widget.NewEntry()
		argsEntry := widget.NewEntry()
	*/

	name := widget.NewFormItem("Name", nameEntry)
	path := widget.NewFormItem("Path", pathEntry)
	args := widget.NewFormItem("Args", argsEntry)
	form := widget.NewForm(name, path, args)

	cancelButton := widget.NewButton("Cancel", func() {
		importWindow.Close()
	})

	okButton := widget.NewButton("OK", func() {
		appName := nameEntry.Text
		appPath := pathEntry.Text
		argsString := argsEntry.Text

		if appName == "" || appPath == "" {
			return
		}

		newApp := executor.App{Name: appName, Path: appPath, Args: argsString}
		importWindow.Close()
		executor.ImportApp(newApp)
		refreshContent()
	})

	nameEntry.SetConfirmButton(okButton)
	pathEntry.SetConfirmButton(okButton)
	argsEntry.SetConfirmButton(okButton)

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
	//counter := 1
	fmt.Println("(re)loading apps")
	apps = executor.GetApps()
	//itemBox := widget.NewVBox()

	itemList := make([]*widget.FormItem, 0)

	for _, app := range *apps {
		appItem := NewAppItem(app)
		formItem := widget.NewFormItem(app.Name, appItem)
		//formItem := widget.NewFormItem(app.Name, appItem)
		//counter++
		itemList = append(itemList, formItem)
		//itemBox.Append(appItem)
	}

	appList = widget.NewForm(itemList...)
	appGroup := widget.NewGroup("Apps", appList)
	appListContainer := fyne.NewContainerWithLayout(layout.NewMaxLayout(), appGroup)
	return appListContainer

	//appList = widget.NewForm(itemList...)
	//fyne.
	//scroller := widget.NewScrollContainer(appList)
	//appListContainer := fyne.NewContainerWithLayout(layout.NewMaxLayout(), scroller)

	//appListScroller := widget.NewScrollContainer(appList)
	//appGroup := widget.NewGroupWithScroller("Apps", appList)
	//appGroup.Resize(fyne.NewSize(400, 330))

	//appListContainer := fyne.NewContainerWithLayout(layout.NewCenterLayout(), appGroup)

	//appListContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), appGroup)
	//appListContainer := widget.NewScrollContainer(appGroup)

	//appScroller := widget.NewScrollContainer(appList)
	//appScroller := fyne.NewContainerWithLayout(layout.NewMaxLayout(), itemBox)
	//appScroller := widget.NewGroupWithScroller("Apps", itemBox)
	//appScroller.Resize(fyne.NewSize(850, 490))
	//appListContainer := fyne.NewContainerWithLayout(layout., appScroller)

}
