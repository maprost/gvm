package internal

import (
	"os"
	"sort"
	"strings"
)

var List Command

func init() {
	List = newCommand("list", "list all installed versions, usage: gvm list [options]")

	List.run = func() {
		ListFunc(homeDir())
	}
}

func ListFunc(home string) {
	Verboseln("home folder: ", home)

	file, err := os.Open(home)
	if err != nil {
		exit("failed opening directory:", home, err)
	}

	list, _ := file.Readdir(0) // 0 to read all files and folders
	names := make([]string, 0, len(list))
	for _, info := range list {
		if info.IsDir() {
			continue
		}

		name := info.Name()
		name = strings.TrimPrefix(name, "go")
		name = strings.TrimSuffix(name, "."+zipType())
		names = append(names, name)
	}
	_ = file.Close()

	sort.Strings(names)
	for _, name := range names {
		Infoln(name)
	}

	Infoln("current go version", currentGoVersion())
}
