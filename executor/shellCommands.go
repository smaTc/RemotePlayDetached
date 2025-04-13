package executor

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var threaded bool = false
var exitAfterExec = false
var isProtonExecution = false

//var currentWorkingDirectory = rpdPath

// Threaded func
func Threaded(b bool) {
	threaded = b
}

// SetExitAfterExec func
func SetExitAfterExec(b bool) {
	exitAfterExec = b
}

// IsThreaded func
func IsThreaded() bool {
	return threaded
}

func executeApp(app App) error {
	p, _ := os.Getwd()
	fmt.Println("current path before execution:", p)
	var argsArray []string

	if app.Args != "" {
		argsArray = strings.Split(app.Args, " ")
	}

	//path, executable, seperator := seperatePathFromExecutable(app.GamePath)
	//os.Chdir(rpdPath)
	//os.Chdir(path)

	var cmd *exec.Cmd

	var gameDirectory, executable, seperator, execLine string
	var sepError error

	if runtime.GOOS == "windows" {
		seperator = "\\"
	} else {
		seperator = "/"
		if strings.Contains(executable, ".exe") {
			isProtonExecution = true
		}
	}

	executable, gameDirectory, sepError = separatePathFromExec(app.GamePath, seperator)

	if sepError != nil {
		return sepError
	}

	if isProtonExecution {
		log.Println("proton execution")

		envs := make([]string, 0)

		prefixPath := app.CompatDataPath + "pfx"
		compatInstallEnv := "STEAM_COMPAT_CLIENT_INSTALL=" + os.Getenv("HOME") + "/.steam"
		compatDataEnv := "STEAM_COMPAT_DATA_PATH=" + app.CompatDataPath
		prefixEnv := "WINEPREFIX=" + prefixPath

		envs = append(envs, compatInstallEnv, compatDataEnv, prefixEnv)
		log.Println("ENVS:", envs)
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, envs...)

		protonExec := "\"" + app.ProtonPath + "\""
		gameExec := "\"" + app.GamePath + "\""
		execLine = protonExec + " run " + gameExec
		os.Chdir(prefixPath)
	} else {
		log.Println("changing dir to path:", gameDirectory)
		os.Chdir(gameDirectory)
		execLine = "." + seperator + executable
	}

	cmd = exec.Command(execLine, argsArray...)

	var err error

	if exitAfterExec {
		err = cmd.Run()
	} else {
		err = cmd.Start()
		go func(err error) {
			cmd.Wait()
			log.Println("error when executing:", err)
		}(err)
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	os.Chdir(rpdPath)
	return nil
}

/*
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
*/

func separatePathFromExec(path, seperator string) (string, string, error) {
	if !strings.Contains(path, seperator) {
		return "", "", errors.New("seperator " + seperator + " not included in path: " + path)
	} else if len(path) == 0 {
		return "", "", errors.New("path cannot be empty")
	}

	splittedPath := strings.Split(path, seperator)

	executable := splittedPath[len(splittedPath)-1]
	directoryPath := strings.Replace(path, executable, "", -1)

	return executable, directoryPath, nil
}
