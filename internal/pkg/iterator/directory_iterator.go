package iterator

import (
	"os"
	"path/filepath"
)

type DirectoryIterator struct {
	path      string
	files     []string
	nextIndex int
}

func NewDirectoryIterator(path string) (*DirectoryIterator, error) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() { // filter out directories
			fileName := dirEntry.Name()
			files = append(files, filepath.Join(path, fileName))
		}
	}

	return &DirectoryIterator{
		path:  path,
		files: files,
	}, nil
}

func (d *DirectoryIterator) HasNext() bool {
	return d.nextIndex < len(d.files)
}

func (d *DirectoryIterator) GetNext() interface{} {
	if d.HasNext() {
		file := d.files[d.nextIndex]
		d.nextIndex++
		return file
	}
	return nil
}
