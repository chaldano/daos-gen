package appdata

import (
    "testing"
    "fmt"
    // "github.com/holmes89/hello-api/translation"
    
)
import conf "trustkbb.de/daosgenerate/appconfig"

func TestPersonRelationsCreate(t *testing.T) {
    // Erzeugt "count" Personen
    // Erzeugt eine Beziehung zwischen Personen
    // Arrange
    path := "./../appconfig/UrlConfig.json"
	conf.Dbconf = conf.InitConfig(path)
	res:=conf.Dbconf.LoadConfig()
	if res != nil {
        t.Errorf(`load Config in %s expected - wrong Config access "%s"`, path,res)
    }
    // Act
	count := 10
    Persons,res1 := CreateNObjects("PersonsNat",count,nil)
	    if res1 != nil {
        t.Errorf(`created person object expected - wrong person object created "%s"`, res1)
    }
	Report(Persons)


    Relations,err := Persons.CreateRelations(1)
    anz := len(Relations.Items)
    fmt.Println("Anzahl Relations",anz)

    // _,err := Persons.CreateRelations(2)
	
    // Assert
    if err != nil {
        t.Errorf(`create person relation expected - wrong relation process "%s"`, res)
    }
    // fmt.Println("Anzeige:",Relations)
    Report(Relations)
}
func TestPersonRelationsByObject(t *testing.T) {
    // Erzeugt "count" Personen
    // Erzeugt eine n * Beziehungen zwischen Personen
    // Arrange
    path := "./../appconfig/UrlConfig.json"
	conf.Dbconf = conf.InitConfig(path)
	res:=conf.Dbconf.LoadConfig()
	if res != nil {
        t.Errorf(`load Config in %s expected - wrong Config access "%s"`, path,res)
    }
    // Act
	count := 10
    n:= 5
    Persons,res1 := CreateNObjects("PersonsNat",count,nil)
	    if res1 != nil {
        t.Errorf(`created person object expected - wrong person object created "%s"`, res1)
    }
	// Report(Persons)
    err := Persons.CreateRelationsByObject(n)
    // Assert
    if err != nil {
        t.Errorf(`create person relation by object expected - wrong relation process "%s"`, res)
    }
    // fmt.Println("Anzeige:",Relations)
    Report(Persons)
}
