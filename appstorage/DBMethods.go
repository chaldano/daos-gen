package appstorage

import (
	"bytes"
	"database/sql"
	"fmt"
	// "log"
	// "strconv"
	// _ "github.com/lib/pq"
)
import errmod "trustkbb.de/daosgenerate/apperror"
import log "github.com/sirupsen/logrus"

// import data "trustkbb.de/daosgenerate/appdata"

var Buf bytes.Buffer

// var LoggerDB *log.Logger

var dbtpg  DBTarget
var dbtora DBTarget

//globale Variable für Dbmanager
var DBM DBManager

func init() {
	log.Info("Setup Database (Init DBMethods)")
		
	dbtpg.name = "postgres"
	dbtpg.host = "att"
	dbtpg.port = 5432
	dbtpg.user = "kbb"
	dbtpg.password = "dime!trans"
	dbtpg.dbname = "trustbusiness"

	dbtora.name = "oracle"
	dbtora.host = "odb"
	dbtora.port = 1521
	dbtora.user = "kbb"
	dbtora.password = "67890"
	dbtora.dbname = "xepdb1"
	
	DBM.DBPool = make([]*DBTarget, 2)
	DBM.DBPool[0] = &dbtpg
	DBM.DBPool[1] = &dbtora
	DBM.Current = &dbtpg

	err := setupDB(DBM.DBPool)
	errmod.CheckError("SetupDB: ", err)
}

// Setup of current Database by name
func(dbm *DBManager) SetCurrentDB(dbname string) error{
	dp := dbm.DBPool;
	for _, dbt := range dp {
		if dbt.name == dbname {
			dbm.Current=dbt;
			fmt.Println("CurrentDB:",dbt.name)
			fmt.Println("CurrentDBLink:",dbt.db)
			return nil;
		}
	}
	return errmod.OopsError("Database "+ dbname + " not found")	
	// return errmod.OopsError("Database not found")		
}

func setupDB(dp []*DBTarget) error {
	// Initialisiere Datenbank
	for _, dbt := range dp {
		// fmt.Println("Name:",dbt.name)
		err := dbt.CreateDBLink()
		if err != nil {
			return err
		}

		if dbt.name != "oracle" {
			// Tabellen-Namen
			sqlT := new(SqlRequest)
			sqlT.SetRequest(dbt, "information_schema.tables", "SELECT table_name", "table_schema = 'public'")
			err = sqlT.DoQuery()
			if err != nil {
				return err
			}
			for _, t := range sqlT.DBResponse.Rows {
				var table Table
				table.name = t.dbvalue[0]
				dbt.tables = append(dbt.tables, table)
			}
			for _, t := range dbt.tables {
				sqlC := new(SqlRequest)
				sqlC.SetRequest(dbt, "information_schema.columns", "SELECT column_name", fmt.Sprintf("table_name = '%s'", t.name))
				err = sqlC.DoQuery()
				if err != nil {
					return err
				}
				for _, c := range sqlC.DBResponse.Rows {
					t.columns = append(t.columns, c.dbvalue[0])
				}
			}
		}
	}
	return nil
}

func (dt *DBTarget) CreateDBLink() error {
	log.Info("Connecting to Database " + dt.dbname +" on Host "+dt.host)
	switch dt.name {
	case "postgres":
		pqInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", dt.host, dt.port, dt.user, dt.password, dt.dbname)
		//fmt.Printf("PostgresInfo: %s\n", pqInfo)

		db, err := sql.Open("postgres", pqInfo)
		if err != nil {
			return err
		}
		// Initialisierung der DB-Instanz
		dt.db = db
		// defer db.Close()

		err = db.Ping()
		//fmt.Println("NameDB-Connect:",dt.name)
		if err != nil {
			return err
		}
	case "oracle":
		oraInfo := fmt.Sprintf("oracle://%s:%s@%s:%d/%s", dt.user, dt.password, dt.host, dt.port, dt.dbname)
		//fmt.Printf("OraInfo: %s\n", oraInfo)
		db, err := sql.Open("oracle", oraInfo)
		if err != nil {
			return err
		}
		dt.db = db
		// defer db.Close()

		err = db.Ping()
		if err != nil {
			return err
		}

	}
	return nil

}

func (sql *SqlRequest) SetRequest(dt *DBTarget, tab string, op string, where string) {
	//fmt.Println("NameDB-Request",dt.name)
	sql.DBT = dt.db
	sql.Table = tab
	sql.Operation = op
	sql.Wherecond = where
}
func (sq *SqlRequest) DoQuery() error {
	// fmt.Println("Start RequestDB")

	var db = sq.DBT
	var table = sq.Table
	request := sq.Operation

	request = fmt.Sprintf("%s FROM %s", request, table)
	if sq.Wherecond != "" {
		request = request + fmt.Sprintf(" WHERE %s;", sq.Wherecond)
	} else {
		request = request + ";"
	}

	// fmt.Printf("Requesting %s \n", request)
	rows, err := db.Query(request)
	if err != nil {
		return err
	}
	columns, _ := rows.Columns()
	if err != nil {
		return err
	}

	count := len(columns)
	sq.Columns = make([]interface{}, count)
	sq.ColumnsPtrs = make([]interface{}, count)

	// Umschreiben Type rows in Type DBResponse
	for rows.Next() {
		for i := range columns {
			sq.ColumnsPtrs[i] = &sq.Columns[i]
		}
		rows.Scan(sq.ColumnsPtrs...)
		var dbrow DBRow
		for i, col := range columns {
			val := sq.Columns[i]
			b, ok := val.([]byte)
			var v interface{}
			if ok {
				v = string(b)
			} else {
				v = val
			}
			dbrow.dbindex = append(dbrow.dbindex, col)
			dbrow.dbvalue = append(dbrow.dbvalue, v.(string))
		}
		sq.DBResponse.Rows = append(sq.DBResponse.Rows, dbrow)
		sq.DBResponse.anz++

	}
	return nil
}
func (sq *SqlRequest) DoInsertRequest(values []string, columns []string) error {
	sq.Operation = "INSERT INTO"

	var db = sq.DBT
	var table = sq.Table
	request := sq.Operation

	var valuestring string
	for i := range values {
		if i < (len(values) - 1) {
			valuestring = valuestring + fmt.Sprintf("'%s',", values[i])
		} else {
			valuestring = valuestring + fmt.Sprintf("'%s'", values[i])
		}
	}
	var columnstring string
	for i := range values {
		if i < (len(values) - 1) {
			columnstring = columnstring + fmt.Sprintf("%s,", columns[i])
		} else {
			columnstring = columnstring + fmt.Sprintf("%s", columns[i])
		}
	}
	request = fmt.Sprintf("%s %s (%s) VALUES (%s)", request, table, columnstring, valuestring)
	log.Debug(fmt.Sprintf("%s\n", request))
	_, err := db.Exec(request)
	if err != nil {
	//errmod.CheckError("SQL-Insert-Operation: ", err)
	return err
	}
	return nil
	// errmod.CheckError("SQL-Insert-Operation: ", err)
}

func (sq *SqlRequest) DoDeleteRequest(condition string) {
	sq.Operation = "DELETE FROM"

	var db = sq.DBT
	var table = sq.Table
	request := sq.Operation
	if len(condition) > 0 {
		request = fmt.Sprintf("%s %s  WHERE %s", request, table, condition)
		fmt.Printf(">> %s\n", request)

	} else {
		request = fmt.Sprintf("%s %s", request, table)
		fmt.Printf(">> ALL records %s\n", request)
	}
	_, err := db.Exec(request)
	errmod.CheckError("Error in SQL-DELETE-Operation: ", err)
}

func (sq *SqlRequest) ShowResults() {
	fmt.Printf("Tabelle: %s Records: %d\n\n", sq.Table, sq.DBResponse.anz)
	for i, row := range sq.DBResponse.Rows {
		fmt.Println("Record:", i)

		for i, column := range row.dbindex {
			if i < len(row.dbindex)-1 {
				fmt.Printf("%-12s :%-15s\n", column, row.dbvalue[i])
			} else {
				fmt.Printf("%-12s :%-15s\n\n", column, row.dbvalue[i])
			}
		}
	}
}
