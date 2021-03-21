package main

import (
	"fmt"
	"github.com/ofcoursedude/wg-manage/cmd/format"
	"os"
	"strings"

	"github.com/ofcoursedude/wg-manage/cmd/add"
	"github.com/ofcoursedude/wg-manage/cmd/bootstrap"
	"github.com/ofcoursedude/wg-manage/cmd/generate"
	"github.com/ofcoursedude/wg-manage/cmd/initialize"
	"github.com/ofcoursedude/wg-manage/cmd/remove"
)

var availableCommands []Command

func main() {

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	availableCommands = []Command{
		add.Add{},
		bootstrap.Bootstrap{},
		generate.Generate{},
		initialize.Initialize{},
		remove.Remove{},
		format.Format{},
	}
	cmd := strings.ToLower(os.Args[1])
	for _, command := range availableCommands {
		if cmd == command.ShortCommand() || cmd == command.LongCommand() {
			command.Run()
			os.Exit(0)
		}
	}
	printHelp()
}

func printHelp() {
	fmt.Println("wg-manage is a command line tool to centrally organize your wireguard configuration.")
	for _, command := range availableCommands {
		command.PrintHelp()
	}
}
