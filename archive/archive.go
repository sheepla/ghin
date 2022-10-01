package archive

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/mholt/archiver/v4"
)

type InternalFileInfo struct {
	Path     string
	FileInfo fs.FileInfo
}

//nolint:gomnd
func IdentifyFormat(fileName string) (archiver.Format, error) {
	//nolint:nosnakecase
	archiveFile, err := os.OpenFile(fileName, os.O_RDONLY, 0o666)
	if err != nil {
		return nil, fmt.Errorf("failed to open the archive file %s: %w", fileName, err)
	}

	format, _, err := archiver.Identify(fileName, archiveFile)
	if err != nil {
		return nil, fmt.Errorf("failed to identify archive format %s: %w", fileName, err)
	}

	return format, nil
}

func ListInternalFiles(fileName string) (*[]InternalFileInfo, error) {
	fsys, err := archiver.FileSystem(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to create virtual file system from the archive file %s: %w",
			fileName,
			err,
		)
	}

	var infos []InternalFileInfo

	if err := fs.WalkDir(fsys, ".", func(path string, ent fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if ent.IsDir() {
			// skip directory
			return nil
		}

		inf, err := ent.Info()
		if err != nil {
			return fmt.Errorf(
				"failed to get directory entry info %s inside the archive file %s: %w",
				path,
				fileName,
				err,
			)
		}

		info := InternalFileInfo{
			Path:     path,
			FileInfo: inf,
		}

		infos = append(infos, info)

		return nil
	}); err != nil {
		return nil, fmt.Errorf(
			"an error occurred while walking inside the archive file %s: %w",
			fileName,
			err,
		)
	}

	return &infos, nil
}
