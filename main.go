package main

import (
	"fmt"
	"os"

	"github.com/smaTc/RemotePlayDetached/executor"
	"github.com/smaTc/RemotePlayDetached/fynegui"
)

//VERSION const
const VERSION = "0.4"

//noGui bool
var noGui bool = false

var runDirectly string = ""
var runFromList string = ""

func main() {

	checkForArgs()
	executor.Init()

	if runDirectly != "" && runFromList == "" {
		err := executor.RunAppWithArgs("direct", runDirectly)
		if err != nil {
			fynegui.TextPopup(err.Error(), "Error:")
		}
	} else if runDirectly == "" && runFromList != "" {
		err := executor.RunAppWithArgs("list", runFromList)
		if err != nil {
			fynegui.TextPopup(err.Error(), "Error:")
		}
	} else if runDirectly != "" && runFromList != "" {
		fmt.Println("you cannot use -r and -g together")
		os.Exit(1)
	} else if noGui && runDirectly == "" && runFromList == "" {
		fmt.Println("Can't run without gui when no game to application to launch is specified")
	}

	if !noGui {
		fynegui.VERSION = VERSION
		fynegui.Run()
	}
}

func checkForArgs() {
	args := os.Args
	for index, arg := range args {
		if arg == "-s" || arg == "-silent" {
			noGui = true
			executor.SetExitAfterExec(true)
			fynegui.SetExitAfterExec(true)
		}

		if arg == "-r" || arg == "-run" {
			runDirectly = args[index+1]
		}

		if arg == "-a" || arg == "-app" {
			runFromList = args[index+1]
		}

		if arg == "-rs" || arg == "-runsilent" {
			noGui = true
			executor.SetExitAfterExec(true)
			fynegui.SetExitAfterExec(true)
			runDirectly = args[index+1]
		}

		if arg == "-as" || arg == "-appsilent" {
			noGui = true
			executor.SetExitAfterExec(true)
			fynegui.SetExitAfterExec(true)
			runFromList = args[index+1]
		}

		if arg == "-h" || arg == "-help" {
			fmt.Println("===============================================================")
			fmt.Println("Possible arguments for Remote Play Detached:")
			fmt.Println("")
			fmt.Println("-s or -silent to disable the gui")
			fmt.Println("-a or -app to run an app from your list by its name")
			fmt.Println("-as or -appsilent to run an app from list without gui")
			fmt.Println("-r or -run to run an app from a given path")
			fmt.Println("-rs or -runsilent to run a game from a given path without gui")
			fmt.Println("")
			fmt.Println("===============================================================")
			os.Exit(0)
		}
	}
}
