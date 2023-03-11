package appstorage

import (
    "testing"
    "fmt"
    // "github.com/holmes89/hello-api/translation"
    
)
import dat "trustkbb.de/tools/dbtools/appdata"

// Kann ein Objekt gespeichert werden
    
func TestPersonsNatDB(t *testing.T) {
    fmt.Println("Durch")

    // Erzeugt "count" Personen
    // Speichert Personen in Datenbank
    // Arrange
	count := 100
    // Act
    Persons,res1 := dat.CreateNObjects("PersonsNat",count,nil)
	if res1 != nil {
        t.Errorf(`created person object expected - wrong person object created "%s"`, res1)
    }
	dat.Report(Persons)
    res := StorePersonsNat(Persons, "oracle")
	// Assert
    if res != nil {
        t.Errorf(`insert person object to DB expected - wrong DB access "%s"`, res)
    }
}

func TestPersonsJurDB(t *testing.T) {
    // Erzeugt "count" Personen
    // Speichert Personen in Datenbank
    // Arrange
    // path := "./../appconfig/UrlConfig.json"
	// conf.Dbconf = conf.InitConfig(path)
	// res:=conf.Dbconf.LoadConfig()
	// if res != nil {
    //     t.Errorf(`load Config in %s expected - wrong Config access "%s"`, path,res)
    // }
    // Act
	count := 10
    Persons,res1 := dat.CreateNObjects("PersonsJur",count,nil)
	if res1 != nil {
        t.Errorf(`created organization object expected - wrong organization object created "%s"`, res1)
    }
	dat.Report(Persons)
    res := StorePersonsJur(Persons, "oracle")
	// Assert
    if res != nil {
        t.Errorf(`insert organization object to DB expected - wrong DB access "%s"`, res)
    }
}
