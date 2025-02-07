package internal

import (
	"fmt"
	"os"
	"sort"
)

var Clear Command

func init() {
	Clear = newCommand("clear", "clear old versions, -n [Int]: how many items should be saved (default: 5), usage: gvm clear [options]")
	n := nFlag(&Clear)

	Clear.run = func() {
		ClearFunc(homeDir(), *n)
	}
}

func ClearFunc(home string, n int) {
	Verboseln("home folder: ", home)

	file, err := os.Open(home)
	if err != nil {
		exit("failed opening directory:", home, err)
	}

	list, _ := file.Readdir(0) // 0 to read all files and folders
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name() < list[j].Name()
	})

	longestName := 0
	for _, info := range list {
		if info.IsDir() {
			continue
		}
		nameLen := len(info.Name())
		if longestName < nameLen {
			longestName = nameLen
		}
	}

	for i, info := range list {
		if info.IsDir() {
			continue
		}

		name := info.Name()
		for j := len(name); j < longestName; j++ {
			name += " "
		}

		mSize := fmt.Sprintf("%dMB", info.Size()/1024/1024)

		action := ""
		if len(list) < i+1+n {
			// do nothing
			action = ""
		} else {
			// clear file
			action = "deleted"
			path := home + "/" + info.Name()
			err = os.Remove(path)
			if err != nil {
				panic(err)
			}
		}

		fmt.Println(name, mSize, "\t", action)
	}
}
