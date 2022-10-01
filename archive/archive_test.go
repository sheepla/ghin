//nolint
package archive

import (
	"path/filepath"
	"testing"

	loader "github.com/peteole/testdata-loader"
)

var (
	baseDir     = loader.GetBasePath()
	archiveFile = filepath.Join(baseDir, "_testdata", "test.zip")
)

func TestIdentifyFormat(t *testing.T) {
	format, err := IdentifyFormat(archiveFile)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("FORMAT:", format.Name())
}

func TestListInternalFiles(t *testing.T) {
	infos, err := ListInternalFiles(archiveFile)
	if err != nil {
		t.Fatal(err)
	}

	for _, info := range *infos {
		t.Logf(
			"%s %d %s %s",
			info.FileInfo.Mode(),
			info.FileInfo.Size(),
			info.FileInfo.ModTime(),
			info.Path,
		)
	}
}
