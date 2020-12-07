package internal

func mainHelp(commands Commands) {
	Infoln("usage: gvm <command> [<args>]")
	Infoln("The commands are:")
	for _, command := range commands {
		Infof("%s\t%s\n", command.name, command.description)
	}

	Infoln("current go version: ", currentGoVersion())
	return
}
