package main

import (
	l4g "code.google.com/p/log4go"
	"database/sql"
	"github.com/docopt/docopt-go"
	"github.com/edmt/geiger/db"
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
	connectionParamaters := db.ConnectionParameters{
		Host:     os.Getenv("GCHOST"),
		Port:     os.Getenv("GCPORT"),
		User:     os.Getenv("GCSQLUSER"),
		Password: os.Getenv("GCSQLPASSWORD"),
		Database: os.Getenv("GCDATABASE"),
	}
	conn := connectionParamaters.MakeConnection()
	defer conn.Close()
	db.Ping(conn)
	c := make(<-chan int)

	if arguments["count"].(bool) {
		c = WriteCount(GenCount(arguments), conn)
	}
	l4g.Info("geiger stopped")
	time.Sleep(time.Second)
	<-c
}

func GenCount(options map[string]interface{}) <-chan GeigerRecord {
	out := make(chan GeigerRecord)
	go func() {
		globPatternList := GetGlobPatternList(options)
		l4g.Info("Directorios encontrados: %d", len(globPatternList))

		for _, globPatternTuple := range globPatternList {
			files, _ := ListFiles(globPatternTuple.Path)
			l4g.Info("%d archivos en directorio %s", len(files), globPatternTuple.Path)
			for _, filePath := range files {
				l4g.Debug("PROCESA: %s", filePath)
				out <- CountFile(TupleRFCFilepath{globPatternTuple, filePath})
			}
		}
		close(out)
	}()
	return out
}

func WriteCount(in <-chan GeigerRecord, conn *sql.DB) <-chan int {
	out := make(chan int)
	go func() {
		for record := range in {
			record.Save(conn)
		}
		out <- 1
		close(out)
	}()
	return out
}
