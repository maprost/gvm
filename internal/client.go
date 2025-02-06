package internal

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dustin/go-humanize"
)

type Client interface {
	GetLatestVersion() (string, bool)
	DownloadFile(filepath string, url string) error
}

type client struct{}

func NewClient() Client {
	return client{}
}

func (client) GetLatestVersion() (string, bool) {
	// get newest go version
	resp, err := http.Get("https://golang.org/VERSION?m=text")
	if err != nil {
		Infoln("Can't get the latest version")
		Verboseln("Can't connect golang.org", err.Error())
		return "", false
	}

	versionBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Infoln("Can't get the latest version")
		Verboseln("Can't open response", err.Error())
		return "", false
	}

	return clearGoVersion(versionBytes, false), true
}

type WriteCounter struct {
	Total        uint64
	HeaderLength string
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)

	fmt.Printf("\rDownloading... %s/%s complete ", humanize.Bytes(wc.Total), wc.HeaderLength)
	return n, nil
}

func (client) DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check if the file
	if resp.ContentLength < 1000000 {
		return errors.New("wrong version")
	}

	tmpFilePath := filepath + ".tmp"
	out, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}

	counter := &WriteCounter{HeaderLength: humanize.Bytes(uint64(resp.ContentLength))}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		_ = out.Close()
		_ = os.Remove(tmpFilePath)
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Println()

	_ = out.Close()

	return os.Rename(tmpFilePath, filepath)
}
