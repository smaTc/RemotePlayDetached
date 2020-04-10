package fynegui

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sort"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/smaTc/RemotePlayDetached/executor"
)

var currentDirectory *widget.Group
var selectedItem string
var currentAppPath *ButtonEntry
var explorerWindow fyne.Window
var currentPath string
var osSeparator = ""

//FileExplorer func
func FileExplorer(pathEntry *ButtonEntry) {
	if runtime.GOOS == "windows" {
		osSeparator = "\\"
	} else {
		osSeparator = "/"
	}

	selectedItem = ""
	currentPath = ""
	currentAppPath = pathEntry
	currentPath = executor.RpdPath()
	openExplorerWindow(currentPath)
}

func openExplorerWindow(path string) {
	explorerWindow = rpd.NewWindow("File Explorer")
	explorerWindow.Resize(fyne.NewSize(600, 500))
	currentDirectory := getDirectoryContent(path)
	buttonColumn := explorerButtonColumn()

	explorerContainer := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, buttonColumn, nil, nil), currentDirectory, buttonColumn)
	explorerWindow.SetContent(explorerContainer)

	explorerWindow.Show()
}

func refreshExplorer() {
	buttonColumn := explorerButtonColumn()
	currentDirectory := getDirectoryContent(currentPath)
	explorerContainer := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, buttonColumn, nil, nil), currentDirectory, buttonColumn)
	explorerWindow.SetContent(explorerContainer)

}

func getDirectoryContent(path string) *widget.Group {
	content, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Could not read directory", path)
	}

	dirs := make([]os.FileInfo, 0)
	files := make([]os.FileInfo, 0)

	for _, element := range content {
		if element.IsDir() {
			dirs = append(dirs, element)
		} else {
			files = append(files, element)
		}
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Name() < dirs[j].Name()
	})

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	directory := make([]*widget.FormItem, 0)

	previousDir := NewClickLabel(".."+osSeparator, func() {
		sepCounter := strings.Count(currentPath, osSeparator)
		pathArray := strings.Split(currentPath, osSeparator)
		fmt.Println("PathArray", pathArray)
		newPath := ""

		for i := 0; i < sepCounter; i++ {
			if i == 0 && sepCounter != 1 {
				newPath += pathArray[i]
			} else {
				newPath += osSeparator + pathArray[i]
			}
		}
		currentPath = newPath
		refreshExplorer()
	})
	prevDirEntry := widget.NewFormItem("", previousDir)

	directory = append(directory, prevDirEntry)

	for _, dir := range dirs {
		dirName := dir.Name()
		entry := NewClickLabel(dirName+osSeparator, func() {
			if currentPath == osSeparator {
				currentPath += dirName
			} else {
				currentPath += osSeparator + dirName
			}
			refreshExplorer()
		})
		formEntry := widget.NewFormItem("", entry)
		directory = append(directory, formEntry)
	}

	for _, file := range files {
		fileName := file.Name()
		entry := NewClickLabel(fileName, func() {
			selectedItem = fileName
			refreshExplorer()
		})
		formEntry := widget.NewFormItem("", entry)
		directory = append(directory, formEntry)
	}

	directoryList := widget.NewForm(directory...)
	directoryGroup := widget.NewGroupWithScroller(path, directoryList)

	return directoryGroup
}

func explorerButtonColumn() *fyne.Container {
	selectButton := widget.NewButton("Select", func() {
		fmt.Println("clicked Select")
		newAppPath := currentPath + osSeparator + selectedItem
		currentAppPath.SetText(newAppPath)
		explorerWindow.Close()
	})

	selectedLabel := widget.NewLabel("Selected: " + selectedItem)

	cancelButton := widget.NewButton("Cancel", func() {
		fmt.Println("clicked Cancel")
		explorerWindow.Close()
	})

	buttonBar := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), selectButton, layout.NewSpacer(), selectedLabel, layout.NewSpacer(), cancelButton)
	return buttonBar

}
