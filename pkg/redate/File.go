package redate

import (
	"os"
)

type File struct {
	OriginalName string
	RedateName   string
}

func CreateFromDirEntry(dirEntry os.DirEntry) *File {
	redateFile := new(File)
	redateFile.OriginalName = dirEntry.Name()
	redateFile.CalculateAndSetNewRedateName()
	return redateFile
}

func (rf *File) CalculateAndSetNewRedateName() {
	rf.RedateName = rf.OriginalName
}
