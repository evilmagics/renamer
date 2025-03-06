package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func rename(oldName, newName string) error {
	err := os.Rename(oldName, newName)
	if err != nil {
		return err
	}

	fmt.Printf("Renamed %s to %s\n", filepath.Base(oldName), filepath.Base(newName))
	return nil
}

func Rename(dir string, options ...Option) Result {
	wg := sync.WaitGroup{}
	res := Result{}
	now := time.Now()

	files, err := getFiles(dir)
	if err != nil {
		fmt.Printf("Error getting files in %s: %v\n", dir, err)
		return res
	}
	res.Count = len(files)

	for _, f := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for _, opt := range options {
				opt(&f)
			}

			err := rename(f.Path, f.AbsPath())
			if err != nil {
				res.Errors = append(res.Errors, FileError{f.Name, err.Error()})
				res.Failed++
				return
			}
			res.Success++
		}()
	}
	wg.Wait()
	res.Elapsed = time.Since(now)

	return res
}

type Result struct {
	Count   int
	Failed  int
	Success int
	Errors  []FileError
	Elapsed time.Duration
}

type FileError struct {
	Name string
	Err  string
}
