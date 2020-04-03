package main

import (
	"fmt"
	"os"

	"github.com/smaTc/RemotePlayDetached/executor"
	"github.com/smaTc/RemotePlayDetached/fynegui"
)

//VERSION const
const VERSION = "0.2"

//noGui bool
var noGui bool = false

var runDirectly string = ""
var runFromList string = ""

func main() {

	checkForArgs()
	executor.Init()

	if runDirectly != "" && runFromList == "" {
		executor.RunAppWithArgs("direct", runDirectly)
	} else if runDirectly == "" && runFromList != "" {
		executor.RunAppWithArgs("list", runFromList)
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
		}

		if arg == "-r" || arg == "-run" {
			runDirectly = args[index+1]
		}

		if arg == "-g" || arg == "-game" {
			runFromList = args[index+1]
		}

		if arg == "-rs" || arg == "-runsilent" {
			noGui = true
			executor.SetExitAfterExec(true)
			runDirectly = args[index+1]
		}

		if arg == "-gs" || arg == "-gamesilent" {
			noGui = true
			executor.SetExitAfterExec(true)
			runFromList = args[index+1]
		}

		if arg == "-h" || arg == "-help" {
			fmt.Println("===============================================================")
			fmt.Println("Possible arguments for Remote Play Detached:")
			fmt.Println("")
			fmt.Println("-s or -silent to disable the gui")
			fmt.Println("-g or -game to run a game from your list by its name")
			fmt.Println("-gs or -gamesilent to run a game from list without gui")
			fmt.Println("-r or -run to run a game with a given path")
			fmt.Println("-rs or -runsilent to run a game with a given path without gui")
			fmt.Println("")
			fmt.Println("===============================================================")
			os.Exit(0)
		}
	}
}
