package appstorage

import (
	"database/sql"
	// "fmt"
	// _ "github.com/lib/pq"
	// "reflect"
)

// import aerr "trustkbb.de/tools/dbtools/apperror"

// DB-Datenstrukturen
// Über den Target entstehen die erforderlichen Methoden.

type DBTarget struct {
	name	 string	// sprechender Bezeichner
	host     string
	port     int
	user     string
	password string
	dbname   string
	db       *sql.DB // wird erst nach Connect initialisiert
	tables	[]Table
}

// DBManager: Pool aus Datenbankverbindungen
type DBManager struct  {
	Current  *DBTarget
	DBPool[] *DBTarget

}

type Table struct{
	name string
	columns []string
}

type DBRow struct {
	dbindex []string
	dbvalue []string
}
type DBRows struct {
	anz  int
	Rows []DBRow
}

type SqlRequest struct {
	DBT         *sql.DB
	Operation   string // Select
	Table       string
	Wherecond   string
	Columns     []interface{}
	ColumnsPtrs []interface{}
	DBResponse  DBRows
	ValueList 	[]string
	ColumnList 	[]string
}

// type DBTables struct {
// }