package appstorage

import (
	"database/sql"
//	_ "github.com/lib/pq"
)
// import data "trustkbb.de/daosgenerate/appdata"
import _ "github.com/lib/pq"
import _ "github.com/sijms/go-ora/v2"



type DBConnect interface {
    CreateDBLink (dbname string)
    GetDBStatus (db *sql.DB)   
    CloseLink (db *sql.DB)
}
