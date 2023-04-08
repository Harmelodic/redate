package io

import (
	"log"
	"os"
	"redate/pkg/redate"
)

func GetRedateFiles(directory string) ([]*redate.File, error) {
	files, err := getFiles(directory)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var redateFiles []*redate.File
	for _, file := range files {
		if !file.IsDir() {
			redateFiles = append(redateFiles, redate.CreateFromDirEntry(file))
		}
	}
	return redateFiles, nil
}

func getFiles(directory string) ([]os.DirEntry, error) {
	currentDirectory, err := os.Open(directory)
	if err != nil {
		log.Fatal(err)
	}
	defer func(currentDirectory *os.File) {
		err := currentDirectory.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(currentDirectory)

	entries, err := currentDirectory.ReadDir(-1)
	if err != nil {
		log.Fatal(err)
	}

	files := filterOutDirectories(entries)
	return files, nil
}

func filterOutDirectories(entries []os.DirEntry) []os.DirEntry {
	var files []os.DirEntry
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry)
		}
	}
	return files
}
