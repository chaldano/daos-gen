package appdata

import (
    "testing"
    // "fmt"
    // "github.com/holmes89/hello-api/translation"
    
)
// import conf "trustkbb.de/tools/dbtools/appconfig"

// Kann ein Objekt gespeichert werden
func TestPersonNatEntity(t *testing.T) {
    
    // Act
	count := 10
    Persons,res := CreateNObjects("PersonsNat",count,nil)
	    
	// Assert
    if res != nil {
        t.Errorf(`created person object expected - wrong person object created "%s"`, res)
    }    
    // fmt.Println("Report:",Persons)
    Report(Persons)
}

func TestPersonJurEntity(t *testing.T) {
    
    // Act
	count := 10
    Persons,res := CreateNObjects("PersonsJur",count,nil)
	    
	// Assert
    if res != nil {
        t.Errorf(`created company object expected - wrong company object created "%s"`, res)
    }
         
    Report(Persons)
}
