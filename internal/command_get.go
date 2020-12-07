package internal

import (
	"github.com/maprost/gvm/internal/osx"
)

var Get Command

func init() {
	Get = newCommand("get", "download a version, usage: gvm get [options] versions")

	Get.run = func() {
		GetFunc(NewClient(), homeDir(), Get.Args())
	}
}

func GetFunc(client Client, home string, versions []string) {
	Verboseln("home folder: ", homeDir())

	for _, version := range versions {
		downloadVersion(client, home, version)
	}
}

func downloadVersion(client Client, home string, version string) (string, bool) {
	var ok bool
	if version, ok = checkVersion(version, client); !ok {
		return "", false
	}

	Infoln("download version: ", version)

	filePath := ZipFilePath(home, version)

	// check if version is already there
	if osx.FileExists(filePath) {
		Infoln("version", version, "already downloaded")
		return version, true
	}

	// check if locked version is already there
	if osx.FileExists(lockedZipFilePath(home, version)) {
		Infoln("version", version, "already downloaded")
		return version, true
	}

	err := client.DownloadFile(filePath, versionUrl(version))
	if err != nil {
		Infoln("can't download version:", version)
		return "", false
	}

	return version, true
}
