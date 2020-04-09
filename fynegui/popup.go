package fynegui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

//TextPopup func
func TextPopup(text, popupType string) {
	var window fyne.Window
	init := Init()
	window = rpd.NewWindow("RPD Notification")

	typeLabel := widget.NewLabel(popupType)
	textLabel := widget.NewLabel(text)
	okButton := widget.NewButton("OK", func() {
		window.Close()
		if exitAfterExec {
			rpd.Quit()
		}
	})

	headerContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), typeLabel, layout.NewSpacer())
	buttonContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), okButton, layout.NewSpacer())
	windowContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), headerContainer, layout.NewSpacer(), textLabel, layout.NewSpacer(), buttonContainer)

	window.SetContent(windowContainer)
	if exitAfterExec {
		window.SetMaster()
		window.ShowAndRun()
	} else {
		if init {
			popupsAfterLoading = append(popupsAfterLoading, window)
		} else {
			window.Show()
		}

	}
}
