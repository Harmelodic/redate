package directory

import (
	"log"
	"os"
)

func GetFiles(directory string) ([]os.DirEntry, error) {
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
