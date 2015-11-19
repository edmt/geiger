package main

import "time"

const (
	TIME_FORMAT_LAYOUT           = "2006 Jan 2 at 3:04pm (MST)"
	TIME_FORMAT_LAYOUT_AS_FOLDER = "2006/01/02"
	PARSE_FORMAT_LAYOUT          = "2006-01-02"
)

func FormatAsFolderPath(t time.Time) string {
	return t.Format(TIME_FORMAT_LAYOUT_AS_FOLDER)
}

func Today() time.Time {
	return time.Now().AddDate(0, 0, -1).Local()
}

func ParseDateOption(option interface{}) time.Time {
	var t time.Time
	if option == nil {
		t = Today()
	} else {
		t, _ = time.Parse(TIME_FORMAT_LAYOUT_AS_FOLDER, option.(string))
	}
	return t
}
