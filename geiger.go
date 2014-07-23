package main

import (
	l4g "code.google.com/p/log4go"
	"github.com/docopt/docopt-go"
	"os"
	"time"
)

const LOG_CONFIGURATION_FILE = "logging-conf.xml"

func init() {
	l4g.LoadConfiguration(LOG_CONFIGURATION_FILE)
}

func main() {
	l4g.Info("Process ID: %d", os.Getpid())
	usage := `geiger.

Usage:
  geiger count --path=<path_location> [--rfc=<rfc>] [--date=<date>]
  geiger -h | --help
  geiger -v | --version

Options:
  -h --help     Show this screen.
  -v --version  Show version.
  --rfc=<rfc>   RFC filter to reduce the search space [default: *].
  --date=<date> Date filter to reduce the search space [default: today]. Format: YYYY-MM-DD (zero-padding!)`

	arguments, _ := docopt.Parse(usage, nil, true, "geiger 0.0.0", false)
	l4g.Debug(arguments)
	if arguments["count"].(bool) {
		count(arguments)
	}
	l4g.Info("geiger stopped")
	time.Sleep(time.Second)
}

func count(options map[string]interface{}) {
	globPatternList := GetGlobPatternList(options["--path"].(string))
	l4g.Info("Directorios encontrados: %d", len(globPatternList))

	for _, globPattern := range globPatternList {
		files, _ := ListFiles(globPattern)
		l4g.Info("%d archivos en directorio %s", len(files), globPattern)
		for _, filePath := range files {
			l4g.Debug("Procesando archivo: %s", filePath)
			parseXml(filePath)
		}
	}
}
