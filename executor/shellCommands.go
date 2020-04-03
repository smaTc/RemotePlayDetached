package executor

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var cwd string
var threaded bool = false
var exitAfterExec = false

//Threaded func
func Threaded(b bool) {
	threaded = b
}

//SetExitAfterExec func
func SetExitAfterExec(b bool) {
	exitAfterExec = b
}

//IsThreaded func
func IsThreaded() bool {
	return threaded
}

func executeApp(app App) {
	p, _ := os.Getwd()
	fmt.Println("current path before execution:", p)
	var argsArray []string

	if app.Args != "" {
		argsArray = strings.Split(app.Args, " ")
	}

	path, executable, seperator := seperatePathFromExecutable(app.Path)
	os.Chdir(cwd)
	os.Chdir(path)

	p2, _ := os.Getwd()
	fmt.Println("path for execution:", p2, "; executable to run:", executable)

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command(executable, argsArray...)
	} else {
		cmd = exec.Command("."+seperator+executable, argsArray...)
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	if exitAfterExec {
		fmt.Println("exiting due to silent mode")
		os.Exit(0)
	}
}

func seperatePathFromExecutable(path string) (string, string, string) {
	var splittedPath []string
	var seperator string
	if strings.Contains(path, "/") {
		splittedPath = strings.Split(path, "/")
		seperator = "/"
	} else if strings.Contains(path, "\\") {
		splittedPath = strings.Split(path, "\\")
		seperator = "\\"
	} else {
		splittedPath = nil
	}

	var executable, directoryPath string

	if splittedPath == nil {
		executable = path
		directoryPath = ""
	} else {
		executable = splittedPath[len(splittedPath)-1]
		directoryPath = strings.Replace(path, executable, "", -1)
	}
	return directoryPath, executable, seperator
}
