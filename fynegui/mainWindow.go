package fynegui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

//VERSION string
var VERSION string

var exitAfterExec = false

var rpd fyne.App = nil
var mainContainer *fyne.Container
var mainWindow fyne.Window
var popupsAfterLoading = make([]fyne.Window, 0)

//Run func
func Run() {
	fmt.Println("Run")
	init := Init()
	fmt.Println("init val:", init)
	mainWindow = rpd.NewWindow("Remote Play Detached")
	mainWindow.Resize(fyne.NewSize(600, 350))
	mainWindow.SetContent(buildMainContent())
	mainWindow.SetMaster()
	mainWindow.Show()
	if len(popupsAfterLoading) > 0 {
		for i := 0; i < len(popupsAfterLoading); i++ {
			popupsAfterLoading[i].Show()
		}
	}
	rpd.Run()
}

//Init function for GUI
func Init() bool {
	if rpd == nil {
		rpd = app.New()

		return true
	}
	return false
}

//SetExitAfterExec func
func SetExitAfterExec(b bool) {
	exitAfterExec = b
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
			scrollContainer := widget.NewScrollContainer(licenseLabel)

			okButton := widget.NewButton("OK", func() {

				aboutWindow.Close()

				refreshContent()
			})

			buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), okButton, layout.NewSpacer())
			paragraphContainer := fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(495, 380)), scrollContainer)
			content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), paragraphContainer, layout.NewSpacer(), buttons)
			aboutWindow.SetContent(content)
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

	versionLabel := widget.NewLabel("v" + VERSION)

	exitButton := widget.NewButton("Exit", func() {
		fmt.Println("clicked: Exit")
		rpd.Quit()
	})

	buttonBar := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), importButton, layout.NewSpacer(), versionLabel, layout.NewSpacer(), exitButton)

	return buttonBar
}
