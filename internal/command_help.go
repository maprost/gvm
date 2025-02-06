package internal

func mainHelp(commands Commands) {
	Infoln("usage: gvm <command> [<args>]")
	Infoln("The commands are:")

	longestName := 0
	for _, command := range commands {
		nameLen := len(command.name)
		if longestName < nameLen {
			longestName = nameLen
		}
	}

	for _, command := range commands {
		name := command.name
		for i := len(name); i < longestName; i++ {
			name += " "
		}

		Infof("\t%s    %s\n", name, command.description)
	}

	Infoln("current go version: ", currentGoVersion())
	return
}
