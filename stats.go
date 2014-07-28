package main

import (
	l4g "code.google.com/p/log4go"
	"database/sql"
	"os"
	"regexp"
	"time"
)

type GeigerRecord struct {
	RFC  string
	Date time.Time
	Name string
	Size int64
	UUID sql.NullString
}

func CountFile(tuple TupleRFCFilepath) GeigerRecord {
	stats, _ := os.Stat(tuple.Filepath)
	expr := regexp.MustCompile("(.{36}).{13}$")
	matches := expr.FindStringSubmatch(stats.Name())

	var uuid sql.NullString
	if len(matches) == 2 {
		uuid = sql.NullString{matches[1], true}
	}
	cfdi := GeigerRecord{tuple.Dir.RFC, tuple.Dir.Date, stats.Name(), stats.Size(), uuid}
	l4g.Debug(cfdi)
	return cfdi
}

func (r GeigerRecord) Save(db *sql.DB) {
	l4g.Debug("Guardando: %s", r)
	stmt, err := db.Prepare("insert into archivos VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		l4g.Error(err)
	}
	result, err := stmt.Exec(r.RFC, r.Date, r.Name, r.Size, r.UUID)
	if err != nil {
		l4g.Error(err)
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		l4g.Error(err)
	}
	l4g.Debug("Rows affected = %d\n", rowCount)
}
