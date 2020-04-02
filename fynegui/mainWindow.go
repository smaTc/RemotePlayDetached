package fynegui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var rpd fyne.App
var mainContainer *fyne.Container
var mainWindow fyne.Window

//Run func
func Run() {
	//initGui()
	rpd = app.New()
	mainWindow = rpd.NewWindow("Remote Play Detached")
	mainWindow.Resize(fyne.NewSize(600, 350))
	mainWindow.SetContent(buildMainContent())
	mainWindow.SetMaster()
	mainWindow.ShowAndRun()
}

func buildMainContent() *fyne.Container {
	mainWindow.SetMainMenu(buildMainMenu())
	mainContainer = fyne.NewContainerWithLayout(layout.NewVBoxLayout(), buildAppListContainer(), layout.NewSpacer(), buildButtonBar())
	return mainContainer
}

func refreshContent() {
	mainWindow.SetContent(buildMainContent())
}

func buildMainMenu() *fyne.MainMenu {

	mainMenu := fyne.NewMainMenu(fyne.NewMenu("Menu",
		fyne.NewMenuItem("About", func() {
			fmt.Println("clicked: About")
			aboutWindow := rpd.NewWindow("About")
			aboutWindow.Resize(fyne.NewSize(500, 400))

			licenseLabel := widget.NewLabel(LICENSE)
			licenseLabel.Resize(fyne.NewSize(450, 350))

			scrollContainer := widget.NewScrollContainer(licenseLabel)
			//scrollContainer.Resize(fyne.NewSize(450, 350))
			//vbox := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), scrollContainer)

			okButton := widget.NewButton("OK", func() {

				aboutWindow.Close()

				refreshContent()
			})

			buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), okButton)

			aboutWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), scrollContainer, layout.NewSpacer(), buttons))

			aboutWindow.Show()

		}),
	))
	return mainMenu
}

func buildButtonBar() *fyne.Container {

	importButton := widget.NewButton("Import", func() {
		fmt.Println("clicked: Import")
		importApp()
	})

	//threadCheck := widget.NewCheck("Multithread (Slower, especially over Wifi)", func(state bool) {
	//	executor.Threaded(state)
	//})
	//threadCheck.SetChecked(executor.IsThreaded())

	exitButton := widget.NewButton("Exit", func() {
		fmt.Println("clicked: Exit")
		rpd.Quit()
	})

	//buttonBar := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), importButton, layout.NewSpacer(), threadCheck, layout.NewSpacer(), exitButton)
	buttonBar := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), importButton, layout.NewSpacer(), layout.NewSpacer(), exitButton)

	return buttonBar
}
