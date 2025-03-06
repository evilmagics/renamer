package main

import (
	"os"
	"path/filepath"
)

type File struct {
	Dir  string
	Path string
	Name string
	Ext  string
}

func (f File) AbsPath() string {
	p, _ := filepath.Abs(filepath.Join(f.Dir, f.Name))
	return p
}

func NewFile(dir string, e os.DirEntry) *File {
	if !e.IsDir() {
		p, _ := filepath.Abs(filepath.Join(dir, e.Name()))
		d, _ := filepath.Abs(dir)

		return &File{
			Dir:  d,
			Path: p,
			Name: e.Name(),
			Ext:  filepath.Ext(e.Name()),
		}
	}

	return nil
}

func getFiles(dir string) ([]File, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	files := []File{}
	for _, e := range entries {
		if !e.IsDir() {
			files = append(
				files,
				*NewFile(dir, e),
			)

		}
	}

	return files, nil
}
