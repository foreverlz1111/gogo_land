package structure

import "time"

type Header struct {
	Title     string
	HintTitle string
	Date      time.Time
	Upload    string
	Edit      string
}
