package main

import (
	l4g "code.google.com/p/log4go"
	"os"
	"path/filepath"
	"time"
)

const (
	TRANSACTION_TYPE_FOLDER = "CFDs_Expedidos"
	FILE_TYPE               = "*.xml"
)

type TupleRFCPath struct {
	RFC  string
	Date time.Time
	Path string
}

type TupleRFCFilepath struct {
	Dir      TupleRFCPath
	Filepath string
}

func ListFiles(globPattern string) (matches []string, err error) {
	return filepath.Glob(globPattern)
}

func GetGlobPatternList(options map[string]interface{}) (output []TupleRFCPath) {
	baseDir := options["--path"].(string)
	rfcOption := options["--rfc"].(string)
	date := options["--date"]
	parsedDate := ParseDateOption(date)
	folder := FormatAsFolderPath(parsedDate)
	l4g.Debug(folder)
	rfcList, _ := getRFCList(baseDir, rfcOption)

	for _, dir := range rfcList {
		rfc := substr(dir, len(baseDir)+1, 13)
		l4g.Debug("Substring: %s %d", rfc, len(rfc))
		t := TupleRFCPath{rfc,
			parsedDate,
			filepath.Join(dir, TRANSACTION_TYPE_FOLDER, folder, FILE_TYPE)}
		output = append(output, t)
	}
	return
}

func getRFCList(baseDir, rfc string) (matches []string, err error) {
	return filepath.Glob(filepath.Join(baseDir, rfc))
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
