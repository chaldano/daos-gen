package appstorage

import (
    "testing"
    // "fmt"
    // "github.com/holmes89/hello-api/translation"
    
)
// import conf "trustkbb.de/tools/dbtools/appconfig"
import dat "trustkbb.de/tools/dbtools/appdata"


func TestPersonRelationsDB(t *testing.T) {
    // Erzeugt "count" Personen
    // Erzeugt eine Beziehung zwischen Personen
    // Arrange
    
    // Act
	count := 10
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

    Relations,err := Persons.CreateRelations(9)
	
    // Assert
    if err != nil {
        t.Errorf(`create person relation expected - wrong relation process "%s"`, res)
    }
    // fmt.Println("Anzeige:",Relations)
    dat.Report(Relations)

    //res = Persons.StorePersons("oracle")
    res = StorePersonsRelations(Relations, "oracle")
	// Assert
    if res != nil {
        t.Errorf(`insert persons relations to DB expected - wrong DB access "%s"`, res)
    }
}
func TestPersonRelationsByObjectDB(t *testing.T) {
    // Erzeugt "count" Personen
    // Erzeugt eine Beziehung zwischen Personen
    // Arrange
    
    // Act
	count := 10
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

    Relations,err := Persons.CreateRelations(9)
	
    // Assert
    if err != nil {
        t.Errorf(`create person relation expected - wrong relation process "%s"`, res)
    }
    // fmt.Println("Anzeige:",Relations)
    dat.Report(Relations)

    //res = Persons.StorePersons("oracle")
    res = StorePersonsRelations(Relations, "oracle")
	// Assert
    if res != nil {
        t.Errorf(`insert persons relations to DB expected - wrong DB access "%s"`, res)
    }
}

