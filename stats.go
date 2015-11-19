package main

import (
	l4g "code.google.com/p/log4go"
	"database/sql"
	"os"
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
	doc := parseXml(tuple.Filepath).(Doc)
	l4g.Debug("This: %s", substr(doc.Fecha, 0, 10))
	date, _ := time.Parse("2006-01-02", substr(doc.Fecha, 0, 10))

	uuid := sql.NullString{doc.Complemento.TimbreFiscalDigital.UUID, true}

	cfdi := GeigerRecord{tuple.Dir.RFC, date, stats.Name(), stats.Size(), uuid}
	return cfdi
}

func (r GeigerRecord) Save(db *sql.DB) {
	l4g.Debug("GUARDA: %s", r.Name)
	stmt, err := db.Prepare("insert into archivos VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		l4g.Error(err)
	}
	_, err = stmt.Exec(r.RFC, r.Date, r.Name, r.Size, r.UUID)
	if err != nil {
		l4g.Error(err)
	}
	// rowCount, err := result.RowsAffected()
	// if err != nil {
	// 	l4g.Error(err)
	// }
	// l4g.Debug("Rows affected = %d\n", rowCount)
}
