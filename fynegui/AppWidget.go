package fynegui

/*
import (
	"fmt"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/smaTc/RemotePlayDetached/executor"
)

type appWidgetRenderer struct {
	label *canvas.Text

	objects   []fyne.CanvasObject
	appWidget *AppWidget
}

func (awr *appWidgetRenderer) MinSize() fyne.Size {
	return fyne.Size{}
}

func (awr *appWidgetRenderer) Layout(size fyne.Size) {

}

func (awr *appWidgetRenderer) Refresh() {

}

func (awr *appWidgetRenderer) BackgroundColor() color.Color {
	return theme.PrimaryColor()

}

func (awr *appWidgetRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{}
}

func (awr *appWidgetRenderer) Destroy() {

}

//AppWidget struct
type AppWidget struct {
	widget.BaseWidget
	Title  *widget.Label
	App    executor.App
	Run    *widget.Button
	Edit   *widget.Button
	Delete *widget.Button
}

//NewAppWidget func
func NewAppWidget(app executor.App) fyne.Widget {
	wg := &AppWidget{App: app}
	wg.ExtendBaseWidget(wg)
	wg.Run = widget.NewButton("Run", func() {
		fmt.Println("Run App")
		executor.RunApp(wg.App)
	})
	wg.Edit = widget.NewButton("Edit", func() {
		fmt.Println("Edit App")
		editApp(app)
	})
	wg.Delete = widget.NewButton("Delete", func() {
		fmt.Println("Delete App")
		executor.DeleteApp(wg.App)
	})
	wg.Title.SetText(app.Name)
	return wg
}

//CreateRenderer func
func (aw *AppWidget) CreateRenderer() fyne.WidgetRenderer {
	aw.ExtendBaseWidget(aw)
	//aw.Run.CreateRenderer()
	//aw.Edit.CreateRenderer()
	//aw.Delete.CreateRenderer()

	return appWidgetRenderer{}
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
}*/
