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

// var appList widget.NewVBox
var appListContainer *fyne.Container
var appListGroup *widget.Group

func importApp() {
	importWindow := rpd.NewWindow("Import App")
	importWindow.Resize(fyne.NewSize(400, 150))

	nameEntry := NewButtonEntry()

	pathEntry := NewButtonEntry()

	argsEntry := NewButtonEntry()

	protonPathEntry := NewButtonEntry()

	winePrefixEntry := NewButtonEntry()

	compatDataPathEntry := NewButtonEntry()

	/*
		nameEntry := widget.NewEntry()
		pathEntry := widget.NewEntry()
		argsEntry := widget.NewEntry()
	*/

	name := widget.NewFormItem("Name", nameEntry)

	path := widget.NewFormItem("Game Path", pathEntry)

	args := widget.NewFormItem("Args", argsEntry)

	proton := widget.NewFormItem("Proton Path", protonPathEntry)

	prefix := widget.NewFormItem("Prefix Path", winePrefixEntry)

	compat := widget.NewFormItem("Compat Data Path", compatDataPathEntry)

	form := widget.NewForm(name, path, args, proton, prefix, compat)
	cancelButton := widget.NewButton("Cancel", func() {
		importWindow.Close()
		if explorerWindow != nil {
			explorerWindow.Close()
		}
	})

	okButton := widget.NewButton("OK", func() {
		appName := nameEntry.Text
		appPath := pathEntry.Text
		argsString := argsEntry.Text
		protonPath := protonPathEntry.Text
		prefixPath := winePrefixEntry.Text
		compatPath := compatDataPathEntry.Text

		if appName == "" || appPath == "" {
			return
		}

		newApp := executor.App{Name: appName, GamePath: appPath, Args: argsString, ProtonPath: protonPath, WinePrefixPath: prefixPath, CompatDataPath: compatPath}
		importWindow.Close()
		executor.ImportApp(newApp)
		refreshMainWindow()
	})

	fileExlporerButton := widget.NewButton("File Explorer", func() {
		FileExplorer(pathEntry)
	})

	nameEntry.SetConfirmButton(okButton)
	pathEntry.SetConfirmButton(okButton)
	argsEntry.SetConfirmButton(okButton)
	protonPathEntry.SetConfirmButton(okButton)
	winePrefixEntry.SetConfirmButton(okButton)
	compatDataPathEntry.SetConfirmButton(okButton)

	buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), okButton, layout.NewSpacer(), fileExlporerButton, layout.NewSpacer(), cancelButton)

	importWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), form, layout.NewSpacer(), buttons))

	importWindow.Show()
}

func buildAppListContainer() *widget.Group {
	appListGroup = loadApps()
	return appListGroup
}

func addToList(app executor.App) {
	appItem := NewAppItem(app)
	formItem := widget.NewFormItem(app.Name, appItem)
	appList.AppendItem(formItem)
}

func loadApps() *widget.Group {
	fmt.Println("(re)loading apps")
	apps = executor.GetApps()

	itemList := make([]*widget.FormItem, 0)

	for _, app := range *apps {
		appItem := NewAppItem(app)
		formItem := widget.NewFormItem(app.Name, appItem)
		itemList = append(itemList, formItem)
	}

	appList = widget.NewForm(itemList...)
	appGroup := widget.NewGroupWithScroller("Apps", appList)
	return appGroup
}
