package executor

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var threaded bool = false
var exitAfterExec = false

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

	path, executable, seperator := seperatePathFromExecutable(app.GamePath)
	os.Chdir(rpdPath)
	os.Chdir(path)

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command(executable, argsArray...)
	} else {

		execLine := ""
		envs := make([]string, 0)

		if strings.Contains(executable, ".exe") {
			compatInstallEnv := "STEAM_COMPAT_CLIENT_INSTALL=$HOME/.steam"
			compatDataEnv := "STEAM_COMPAT_DATA_PATH=" + app.CompatDataPath
			prefixEnv := "WINEPREFIX=" + app.CompatDataPath + "pfx"

			envs = append(envs, compatInstallEnv, compatDataEnv, prefixEnv)

			//FIXME: path to proton executable not running due to whitespaces
			protonExec := "\"" + app.ProtonPath + "\""
			execLine = /* compatInstallEnv + " " + compatDataEnv + " " + prefixEnv + " " +  */ protonExec // + " run "

			os.Chdir(app.CompatDataPath + "pfx")
		}

		seperator = seperator
		/* if path != "" {
			execLine += "." + seperator + executable
		} else {
			execLine += executable
		} */

		//execLine += "\"" + app.GamePath + "\""
		log.Println("execLine: ", execLine)
		log.Println("ENVS:", envs)
		//cmd = exec.Command(execLine, argsArray...)
		cmd = exec.Command(app.ProtonPath, []string{"run", app.GamePath}...)
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, envs...)
	}

	p2, _ := os.Getwd()
	fmt.Println("path for execution:", p2, "; executable to run:", executable)

	var err error

	if exitAfterExec {
		err = cmd.Run()
	} else {
		err = cmd.Start()
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
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
