package internal_test

import (
	"os"
	"testing"

	"github.com/maprost/gvm/internal"
	"github.com/maprost/gvm/internal/osx"
	"github.com/maprost/testbox/must"
)

func TestClient_GetLatestVersion(t *testing.T) {
	t.Run("get latest version", func(t *testing.T) {
		c := internal.NewClient()

		v, ok := c.GetLatestVersion()
		must.BeTrue(t, ok)
		must.NotBeEmpty(t, v)
	})
}

func TestClient_DownloadFile(t *testing.T) {
	t.Run("download file from correct url", func(t *testing.T) {
		const file = "./correct_url_1MB.zip"

		c := internal.NewClient()
		err := c.DownloadFile(file, "http://speedtest.ftp.otenet.gr/files/test1Mb.db")
		must.BeNoError(t, err)

		must.BeTrue(t, osx.FileExists(file))
		must.BeNoError(t, os.Remove(file))
	})

	t.Run("download file from incorrect url", func(t *testing.T) {
		const file = "incorrect_url"

		c := internal.NewClient()
		err := c.DownloadFile(file, "http.//wrong-url.de")
		must.BeError(t, err)
		must.BeFalse(t, osx.FileExists(file))
	})

	t.Run("download file to small", func(t *testing.T) {
		const file = "./correct_url_100K.zip"

		c := internal.NewClient()
		err := c.DownloadFile(file, "http://speedtest.ftp.otenet.gr/files/test100k.db")
		must.BeError(t, err)
		must.BeFalse(t, osx.FileExists(file))
	})
}
