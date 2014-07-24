package main

import (
	l4g "code.google.com/p/log4go"
	"os"
	"path/filepath"
)

const (
	TRANSACTION_TYPE_FOLDER = "CFDs_Expedidos"
	FILE_TYPE               = "*.xml"
)

func ListFiles(globPattern string) (matches []string, err error) {
	return filepath.Glob(globPattern)
}

func GetGlobPatternList(options map[string]interface{}) (output []string) {
	baseDir := options["--path"].(string)
	rfc := options["--rfc"].(string)
	date := options["--date"]
	folder := FormatAsFolderPath(ParseDateOption(date))
	l4g.Debug(folder)
	rfcList, _ := getRFCList(baseDir, rfc)

	for _, value := range rfcList {
		output = append(output,
			filepath.Join(value, TRANSACTION_TYPE_FOLDER, folder, FILE_TYPE))
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
