package internal

import (
	"os"
	"path/filepath"

	"github.com/maprost/gvm/internal/osx"
	"github.com/mholt/archiver/v3"
)

var Install Command

func init() {
	Install = newCommand("install", "install a version (eg: 1.17.8, latest), usage: gvm install [options] versions")
	root := rootFlag(&Install)

	Install.run = func() {
		if len(Install.Args()) != 1 {
			exit("need a version to switch")
		}

		InstallFunc(*root, homeDir(), Install.Args()[0], NewClient())
	}
}

func InstallFunc(root string, home string, version string, client Client) {
	Verboseln("home folder: ", home)
	Verboseln("root folder: ", root)

	var ok bool
	if version, ok = downloadVersion(client, home, version); !ok {
		exit()
	}

	Infoln("install version: ", version)
	Infoln("old go version", currentGoVersion())

	removeTmpGo(home)
	defer removeTmpGo(home)

	err := archiver.Unarchive(ZipFilePath(home, version), home)
	if err != nil {
		exit("can't decompress file", ZipFilePath(home, version))
	}

	ok = clearGoRoot(root)
	if !ok {
		exit("can't clear", root)
	}

	err = osx.CopyFolder(filepath.Join(home, "go"), root)
	if err != nil {
		exit("can't copy go files", err.Error())
	}

	//// delete the file
	//Run("rm", file)
	//if err != nil {
	//	fmt.Println("wrong version: ", version)
	//	continue
	//}
	//
	//// rename go folder
	//Run("mv", home+"/go/", goFolder)

	//goFolder := goFolderPath(home, version)
	//
	//// check if version is already there
	//_, err := Output("ls", goFolder)
	//if err != nil {
	//	Infoln("version is not installed. Do: gvm install", version)
	//	os.Exit(1)
	//}
	//
	//Infoln("old ", SimpleOutput("go", "version"))
	//
	//// delete old version
	//Run("sudo", "rm", "-R", root)
	//
	//// set new version
	//Run("sudo", "cp", "-r", goFolder, root)

	Infoln("new go version", currentGoVersion())
}

func removeTmpGo(home string) {
	tmpGoPath := filepath.Join(home, "go")

	err := os.Chdir(tmpGoPath)
	if err != nil {
		return
	}

	// delete folder
	err = os.RemoveAll(tmpGoPath)
	if err != nil {
		exit("Can't delete", tmpGoPath, err.Error())
	}
}

func clearGoRoot(root string) bool {
	file, err := os.Open(root)
	if err != nil {
		Infoln("failed opening directory:", root, err)
		return false
	}

	// delete folders
	list, _ := file.Readdir(0) // 0 to read all files and folders
	ok := true
	for _, info := range list {
		if info.IsDir() {
			p := filepath.Join(root, info.Name())
			err = os.RemoveAll(p)
			if err != nil {
				Infoln("Can't delete folder", p, err.Error())
				ok = false
			}
		} else {
			p := filepath.Join(root, info.Name())
			err = os.Remove(p)
			if err != nil {
				Infoln("Can't delete file", p, err.Error())
				ok = false
			}
		}
	}

	return ok
}
