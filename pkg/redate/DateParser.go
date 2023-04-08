package redate

import "time"

type DateFoundRecord struct {
	Found               bool
	Source              string
	Date                time.Time
	DateString          string
	IndexLocationOfDate int64
	UndatedString       string
}

func findDate(stringContainingDate string) DateFoundRecord {
	var dateFoundRecord = DateFoundRecord{
		Found:               false,
		Source:              stringContainingDate,
		Date:                time.Now(),
		DateString:          "",
		IndexLocationOfDate: 0,
		UndatedString:       "",
	}
	return dateFoundRecord
}
