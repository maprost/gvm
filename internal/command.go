package internal

import (
	"flag"
)

type Command struct {
	*flag.FlagSet
	name        string
	description string
	run         func()

	verbose *bool
}

func newCommand(name string, description string) Command {
	c := Command{
		FlagSet:     flag.NewFlagSet(name, flag.ExitOnError),
		name:        name,
		description: description,
		run:         nil,
	}

	verboseFlag(&c)
	return c
}

type Commands []Command

func (commands Commands) ParseAndRun(args []string) {
	if len(args) == 1 {
		mainHelp(commands)
		return
	}

	// parse subcommands
	commandNotFound := true
	commandToSearchFor := args[1]
	for i := 0; i < len(commands) && commandNotFound; i++ {
		if commandToSearchFor == commands[i].name {
			commands[i].Parse(args[2:])
			verbose = commands[i].verbose

			commands[i].run()
			commandNotFound = false
		}
	}
	if commandNotFound {
		Infof("%q is not valid command.\n", commandToSearchFor)
		mainHelp(commands)
	}
}
