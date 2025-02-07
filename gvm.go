package main

import (
	"os"

	"github.com/maprost/gvm/internal"
)

func main() {
	commands := internal.Commands{
		internal.Get,
		internal.Install,
		internal.List,
		internal.Clear,
	}
	commands.ParseAndRun(os.Args)
}
