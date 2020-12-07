package internal_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/maprost/gvm/internal"
	"github.com/maprost/testbox/must"
	"github.com/maprost/testbox/should"
	"github.com/mholt/archiver/v3"
)

type testClient struct {
	getLatestVersion      string
	getLatestVersionOk    bool
	getLatestVersionCount int
	downloadFileError     error
	downloadFileCount     int
}

func (c *testClient) GetLatestVersion() (string, bool) {
	c.getLatestVersionCount++
	return c.getLatestVersion, c.getLatestVersionOk
}

func (c *testClient) DownloadFile(filepath string, url string) error {
	c.downloadFileCount++
	return c.downloadFileError
}

func testDir() string {
	return os.TempDir()
}

func TestGetFunc(t *testing.T) {
	const version = "1.13.7"

	t.Run("check simple get", func(t *testing.T) {
		c := &testClient{
			getLatestVersionCount: 0,
			downloadFileCount:     0,
		}
		internal.GetFunc(c, testDir(), []string{version})

		should.BeEqual(t, c.downloadFileCount, 1)     // there is a download
		should.BeEqual(t, c.getLatestVersionCount, 0) // no 'latest' check
	})

	t.Run("check if version exists", func(t *testing.T) {
		rawFile := filepath.Join(testDir(), internal.GoFileName(version))

		// create file
		must.BeNoError(t, ioutil.WriteFile(rawFile, []byte("...."), os.ModePerm))

		archiveFile := internal.ZipFilePath(testDir(), version)
		must.BeNoError(t, archiver.Archive([]string{rawFile}, archiveFile))

		c := &testClient{
			getLatestVersionCount: 0,
			downloadFileCount:     0,
		}
		internal.GetFunc(c, testDir(), []string{version})

		should.BeEqual(t, c.downloadFileCount, 0)     // no download
		should.BeEqual(t, c.getLatestVersionCount, 0) // no 'latest' check

		// remove file
		must.BeNoError(t, os.Remove(archiveFile))
		must.BeNoError(t, os.Remove(rawFile))
	})

	t.Run("check if locked version exists", func(t *testing.T) {

	})

	t.Run("check wrong version", func(t *testing.T) {
		c := &testClient{
			getLatestVersionCount: 0,
			downloadFileCount:     0,
			downloadFileError:     fmt.Errorf("not found"),
		}
		internal.GetFunc(c, testDir(), []string{"1.wrong"})

		should.BeEqual(t, c.downloadFileCount, 1)     // there is a download trial
		should.BeEqual(t, c.getLatestVersionCount, 0) // no 'latest' check

	})

	t.Run("check latest version", func(t *testing.T) {
		c := &testClient{
			getLatestVersion:      version,
			getLatestVersionOk:    true,
			getLatestVersionCount: 0,
			downloadFileCount:     0,
		}
		internal.GetFunc(c, testDir(), []string{"latest"})

		should.BeEqual(t, c.downloadFileCount, 1)     // there is a download
		should.BeEqual(t, c.getLatestVersionCount, 1) // there is a 'latest' check
	})
}
