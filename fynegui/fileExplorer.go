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
	"github.com/shirou/gopsutil/disk"
	"github.com/smaTc/RemotePlayDetached/executor"
)

var currentDirectory *widget.Group
var selectedItem string
var currentAppPath *ButtonEntry
var explorerWindow fyne.Window = nil
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
	currentDirectory := getDirectoryList(path)

	rootSel := rootSelector()
	buttonColumn := explorerButtonColumn()

	explorerContainer := fyne.NewContainerWithLayout(layout.NewBorderLayout(rootSel, buttonColumn, nil, nil), rootSel, currentDirectory, buttonColumn)
	explorerWindow.SetContent(explorerContainer)

	explorerWindow.Show()
}

func refreshExplorer() {
	rootSel := rootSelector()
	buttonColumn := explorerButtonColumn()
	currentDirectory := getDirectoryList(currentPath)
	explorerContainer := fyne.NewContainerWithLayout(layout.NewBorderLayout(rootSel, buttonColumn, nil, nil), rootSel, currentDirectory, buttonColumn)
	explorerWindow.SetContent(explorerContainer)

}

func getDirectoryContent(path string) ([]os.FileInfo, []os.FileInfo) {
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

	return dirs, files
}

func getDirectoryList(path string) *widget.Group {

	dirs, files := getDirectoryContent(path)

	directory := make([]*widget.FormItem, 0)

	previousDir := NewClickLabel(".."+osSeparator+" (parent directory)", func() {
		sepCounter := strings.Count(currentPath, osSeparator)
		pathArray := strings.Split(currentPath, osSeparator)
		newPath := ""

		for i := 0; i < sepCounter; i++ {

			if runtime.GOOS == "windows" {
				if i == 0 {
					newPath += pathArray[i] + osSeparator
				} else if i == 1 {
					newPath += pathArray[i]
				} else {
					newPath += osSeparator + pathArray[i]
				}
			} else {
				if i == 0 && sepCounter != 1 {
					newPath += pathArray[i]
				} else {
					newPath += osSeparator + pathArray[i]
				}
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

			if runtime.GOOS == "windows" {
				currentPath = strings.Replace(currentPath, "\\\\", "\\", -1)
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

func rootSelector() *fyne.Container {
	var rootSelect *widget.Select
	if runtime.GOOS == "windows" {
		partitions, _ := disk.Partitions(false)
		drives := make([]string, len(partitions))

		for i := 0; i < len(drives); i++ {
			drives[i] = partitions[i].Mountpoint + osSeparator
		}

		rootSelect = widget.NewSelect(drives, func(newDir string) {
			currentPath = newDir
			refreshExplorer()
		})

	} else {
		dirs, _ := getDirectoryContent("/")
		var dirStrings = make([]string, len(dirs)+1)
		dirStrings[0] = osSeparator
		for i := 0; i < len(dirs); i++ {
			dirStrings[i+1] = osSeparator + dirs[i].Name()
		}

		rootSelect = widget.NewSelect(dirStrings, func(newDir string) {
			currentPath = newDir
			refreshExplorer()
		})
	}
	rootSelect.PlaceHolder = "Drive or Root Directory"

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), rootSelect, layout.NewSpacer())
}

func explorerButtonColumn() *fyne.Container {
	selectButton := widget.NewButton("Select", func() {
		newAppPath := currentPath + osSeparator + selectedItem
		currentAppPath.SetText(newAppPath)
		explorerWindow.Close()
	})

	selectedLabel := widget.NewLabel("Selected: " + selectedItem)

	cancelButton := widget.NewButton("Cancel", func() {
		explorerWindow.Close()
	})

	buttonBar := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), selectButton, layout.NewSpacer(), selectedLabel, layout.NewSpacer(), cancelButton)
	return buttonBar
}
