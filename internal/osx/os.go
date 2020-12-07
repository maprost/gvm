package osx

import (
	"io"
	"os"
	"path/filepath"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	inInfo, err := in.Stat()
	if err != nil {
		return err
	}

	out, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, inInfo.Mode())
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}

func CopyFolder(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	err = os.Chdir(dst)
	if err != nil {
		// if not exists
		err = os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			return err
		}
	}

	inList, _ := in.Readdir(0) // 0 to read all files and folders
	for _, inInfo := range inList {
		inF := filepath.Join(src, inInfo.Name())
		outF := filepath.Join(dst, inInfo.Name())

		var err error
		if inInfo.IsDir() {
			err = CopyFolder(inF, outF)
		} else {
			err = CopyFile(inF, outF)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
