package fynegui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/smaTc/RemotePlayDetached/executor"
)

//AppItem struct
type AppItem struct {
	App executor.App
	widget.Box
}

//NewAppItem func
func NewAppItem(app executor.App) fyne.Widget {
	item := &AppItem{App: app}
	item.ExtendBaseWidget(item)
	item.Box = *widget.NewHBox()

	//item.Box.Append(layout.NewSpacer())
	//seperatorLabel := widget.NewLabel("App: " + app.Name)
	//item.Box.Append(seperatorLabel)
	item.Box.Append(layout.NewSpacer())

	pathLabel := widget.NewLabel("Path: " + app.Path)
	item.Box.Append(pathLabel)
	item.Box.Append(layout.NewSpacer())

	runButton := widget.NewButton("Run", func() {
		fmt.Println("Run App")
		executor.RunApp(app)
	})
	item.Box.Append(runButton)

	editButton := widget.NewButton("Edit", func() {
		fmt.Println("Edit App")
		editApp(app)
	})
	item.Box.Append(editButton)

	deleteButton := widget.NewButton("Delete", func() {
		fmt.Println("Delete App")
		executor.DeleteApp(app)
		refreshContent()
	})
	item.Box.Append(deleteButton)

	return item
}

func editApp(oldApp executor.App) {
	editWindow := rpd.NewWindow("Edit App")
	editWindow.Resize(fyne.NewSize(400, 150))

	nameWidget := widget.NewEntry()
	nameWidget.SetText(oldApp.Name)

	pathWidget := widget.NewEntry()
	pathWidget.SetText(oldApp.Path)

	argsWidget := widget.NewEntry()
	argsWidget.SetText(oldApp.Args)

	name := widget.NewFormItem("Name", nameWidget)
	path := widget.NewFormItem("Path", pathWidget)
	args := widget.NewFormItem("Args", argsWidget)
	form := widget.NewForm(name, path, args)

	cancelButton := widget.NewButton("Cancel", func() {
		editWindow.Close()
	})

	okButton := widget.NewButton("OK", func() {
		appName := name.Widget.(*widget.Entry).Text
		appPath := path.Widget.(*widget.Entry).Text
		argsString := args.Widget.(*widget.Entry).Text

		if appName == "" || appPath == "" {
			return
		}

		newApp := executor.App{Name: appName, Path: appPath, Args: argsString}
		editWindow.Close()
		executor.EditApp(oldApp, newApp)
		refreshContent()
	})

	buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), okButton, layout.NewSpacer(), cancelButton)

	editWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), form, layout.NewSpacer(), buttons))

	editWindow.Show()
}
