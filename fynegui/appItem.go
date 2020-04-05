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
		err := executor.RunApp(app)
		if err != nil {
			TextPopup(err.Error(), "Error:")
		}
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

	nameEntry := NewButtonEntry()
	nameEntry.SetText(oldApp.Name)

	pathEntry := NewButtonEntry()
	pathEntry.SetText(oldApp.Path)

	argsEntry := NewButtonEntry()
	argsEntry.SetText(oldApp.Args)

	name := widget.NewFormItem("Name", nameEntry)
	path := widget.NewFormItem("Path", pathEntry)
	args := widget.NewFormItem("Args", argsEntry)
	form := widget.NewForm(name, path, args)

	cancelButton := widget.NewButton("Cancel", func() {
		editWindow.Close()
	})

	okButton := widget.NewButton("OK", func() {
		appName := nameEntry.Text
		appPath := pathEntry.Text
		argsString := argsEntry.Text

		if appName == "" || appPath == "" {
			return
		}

		newApp := executor.App{Name: appName, Path: appPath, Args: argsString}
		editWindow.Close()
		executor.EditApp(oldApp, newApp)
		refreshContent()
	})

	nameEntry.SetConfirmButton(okButton)
	pathEntry.SetConfirmButton(okButton)
	argsEntry.SetConfirmButton(okButton)

	buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), okButton, layout.NewSpacer(), cancelButton)

	editWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), form, layout.NewSpacer(), buttons))

	editWindow.Show()
}
