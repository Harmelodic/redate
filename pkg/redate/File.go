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

func (f *File) CalculateAndSetNewRedateName() {
	f.CalculateAndSetNewRedateNameFormatted("2006-01-02")
}

func (f *File) CalculateAndSetNewRedateNameFormatted(layout string) {
	record := findDate(f.OriginalName)
	if record.Found {
		f.RedateName = record.Date.Format(layout) + "-" + record.UndatedString
	} else {
		f.RedateName = f.OriginalName
	}
}
