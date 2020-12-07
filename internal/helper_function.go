package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func homeDir() string {
	path, err := os.UserHomeDir()
	if err != nil {
		exit("can't open home directory")
	}

	path = filepath.Join(path, ".gvm")
	err = os.Chdir(path)
	if err != nil {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			exit("can't create", path)
		}
	}

	return path
}

func rootDir() string {
	//root := os.Getenv("GOROOT")
	//if root == "" {
	//	exit("environment variable GOROOT is not set")
	//}
	//return root

	return runtime.GOROOT()
}

func Infoln(m ...interface{}) {
	log.Println(m...)
}

func Infof(msg string, p ...interface{}) {
	log.Printf(msg, p...)
}

var verbose *bool

func Verboseln(m ...interface{}) {
	if verbose != nil && *verbose {
		Infoln(m...)
	}
}

func Verbosef(msg string, p ...interface{}) {
	if verbose != nil && *verbose {
		Infof(msg, p...)
	}
}

func exit(msg ...interface{}) {
	Infoln(msg...)
	os.Exit(1)
}

func checkVersion(version string, client Client) (string, bool) {
	ok := true
	if version == "latest" {
		version, ok = client.GetLatestVersion()
	}

	return version, ok
}

func currentGoVersion() string {
	root := rootDir()
	versionPath := filepath.Join(root, "VERSION")
	version, err := ioutil.ReadFile(versionPath)
	if err != nil {
		Verboseln("Can't open ", versionPath)
		return "unknown"
	}

	return string(version)[2:] // cut 'go' prefix
}

func getArch() string {
	return runtime.GOARCH
}

func getOs() string {
	return runtime.GOOS
}

func zipType() string {
	var t string
	switch getOs() {
	case "linux", "darwin", "freebsd":
		t = "tar.gz"
	case "windows":
		t = "zip"
	default:
		t = "unknown"
	}
	return t
}

func GoFileName(version string) string {
	return fmt.Sprintf("go%s", version)
}

func ZipFilePath(home, version string) string {
	return filepath.Join(home, fmt.Sprintf("%s.%s", GoFileName(version), zipType()))
}

func lockedZipFilePath(home, version string) string {
	return ZipFilePath(home, version+"_locked")
}

func versionUrl(version string) string {
	return fmt.Sprintf("https://dl.google.com/go/go%s.%s-%s.%s", version, getOs(), getArch(), zipType())
}
