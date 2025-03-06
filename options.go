package main

import (
	"path/filepath"
	"strings"
)

type Option func(*File)

func WithPrefix(prefix string) Option {
	return func(f *File) {
		f.Name = prefix + f.Name
	}
}
func WithSuffix(suffix string) Option {
	return func(f *File) {
		ext := filepath.Ext(f.Name)
		f.Name = strings.TrimSuffix(f.Name, ext) + suffix + ext
	}
}
