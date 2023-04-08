package main

import (
	"fmt"
	"redate/pkg/io"
	"redate/pkg/redate"
)

func main() {
	files, err := io.GetFiles("./testFiles")
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
