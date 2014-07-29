package db

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

type ConnectionParameters struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// Construye la cadena conexión
func (cp ConnectionParameters) makeConnectionString() string {
	return fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s;log=2;encrypt=disable",
		cp.Host, cp.Port, cp.User, cp.Password, cp.Database)
}

func (cp ConnectionParameters) MakeConnection() *sql.DB {
	stringConnection := cp.makeConnectionString()
	log.Print(stringConnection)
	db, err := sql.Open("mssql", stringConnection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Revisa que pueda acceder a la base de datos
func Ping(db *sql.DB) {
	err := db.Ping()
	if err == nil {
		log.Print("Everything is ok")
	} else {
		log.Panic(err)
	}
}
