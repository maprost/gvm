package internal

var Clear Command

func init() {
	Clear = newCommand("clear", "clear old versions, usage: gvm clear [options]")
	n := nFlag(&Clear)

	Clear.run = func() {
		ClearFunc(homeDir(), *n)
	}
}

func ClearFunc(home string, n int) {
	Verboseln("home folder: ", home)
}
