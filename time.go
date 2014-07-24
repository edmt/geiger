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

func ParseDateOption(option string) time.Time {
	t, _ := time.Parse(PARSE_FORMAT_LAYOUT, option)
	return t
}
