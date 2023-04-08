package main

import (
	"fmt"
	"redate/src/redate"
)

func main() {
	files, err := getFiles("./testFiles")
	if err != nil {
		return
	}

	var redateFiles []*redate.File
	for _, file := range files {
		if !file.IsDir() {
			redateFiles = append(redateFiles, redate.CreateFromDirEntry(file))
		}
	}

	fmt.Println("Suggestions:")
	logRedateFiles(redateFiles)
}

func logRedateFiles(redateFiles []*redate.File) {
	for _, redateFile := range redateFiles {
		fmt.Println("\t-", redateFile.OriginalName, "->", redateFile.RedateName)
	}
}
