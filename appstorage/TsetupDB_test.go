package appstorage

import (
    "testing"
	"fmt"

    // "github.com/holmes89/hello-api/translation"
)

// Kann auf eine verfügbare DB zugegriffen werden
// Setup Postgres DB
// func TestSetupDBPostgres(t *testing.T) {
//     // Arrange
// 	// var DBT DBTarget
// 	var dbm DBManager
// 	var dbtpg DBTarget
	
// 	dbtpg.name = "postgres"
// 	dbtpg.host = "att"
// 	dbtpg.port = 5432
// 	dbtpg.user = "kbb"
// 	dbtpg.password = "dime!trans"
// 	dbtpg.dbname = "trustbusiness"
	
// 	dbm.DBPool = make([]DBTarget, 1)	
// 	dbm.DBPool[0] = dbtpg
// 	dbm.current = dbtpg

//     // Act
//     res := setupDB(dbm.DBPool)

//     // Assert
//     if res != nil {
//         t.Errorf(`expected "running PostgresDB" but received "%s"`, res)
//     }
// }

func TestSetupDBPostgres(t *testing.T) {
    // Act
	dbt := DBM.Current;

    // Assert
    if dbt.name != "postgres" {
        t.Errorf(`expected "initialized PostgresDB" but received "%s"`, dbt.name)
    }
}
//Setze Postgres als aktuelle DB
func TestSetDBpg(t *testing.T) {
    // Act
	
	db := &DBM
	err:=db.SetCurrentDB("oracle")
	
	// Assert
	if err != nil {
        t.Errorf(`expected "selected OracleDB" but received "%s"`, err)
    }
	dbt := DBM.Current
	fmt.Println("CurrentDB:",dbt.name)
	
}
//Setze Oracle als aktuelle DB
func TestSetDBora(t *testing.T) {
    // Act
	db := &DBM
	err:=db.SetCurrentDB("oracle")
	// Assert
	if err != nil {
        t.Errorf(`expected "selected OracleDB" but received "%s"`, err)
    }
	dbt := DBM.Current
	fmt.Println("CurrentDB:",dbt.name)
	
}

//setup Oracle DB
// func TestSetupDBOracle(t *testing.T) {
//     // Arrange
// 	// var DBT DBTarget
// 	var dbm DBManager
// 	var dbtora DBTarget
	
	
// 	dbtora.name = "oracle"
// 	dbtora.host = "odb"
// 	dbtora.port = 1521
// 	dbtora.user = "kbb"
// 	dbtora.password = "12345"
// 	dbtora.dbname = "xepdb1"
	
// 	dbm.DBPool = make([]DBTarget, 1)	
// 	dbm.DBPool[0] = dbtora
// 	dbm.Current = dbtora

//     // Act
//     res := setupDB(dbm.DBPool)

//     // Assert
//     if res != nil {
//         t.Errorf(`expected "running OracleDB" but received "%s"`, res)
//     }
// }

// // Wird ein falches DB Setup erkannt
// func TestWrongSetupDB(t *testing.T) {
//     // Arrange
// 	// var DBT DBTarget
// 	var dbtarget = map[string]string{
// 		"host": "attx",
// 		"port":  "5432",
// 		"user": "kbb",
// 		"password": "dime!trans",
// 		"dbname": "trustbusiness",
// 	}
	
//     // Act
//     res := setupDB(dbtarget)

//     // Assert
//     if res == nil {
//         t.Errorf(`expected "not-running DB" - wrong server "%s"`, res)
//     }
// }



