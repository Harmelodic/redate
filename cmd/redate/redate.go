package main

import (
	"fmt"
	"redate/pkg/io"
	"redate/pkg/redate"
)

func main() {
	redateFiles, err := io.GetRedateFiles("./test/data")
	if err != nil {
		return
	}

	fmt.Println("Suggestions:")
	logRedateFiles(redateFiles)

	// Ask user for confirmation

	io.RenameFiles(redateFiles)
	fmt.Println("Done")
}

func logRedateFiles(redateFiles []*redate.File) {
	for _, redateFile := range redateFiles {
		fmt.Println("\t-", redateFile.OriginalName, "->", redateFile.RedateName)
	}
}
