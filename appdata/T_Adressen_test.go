package appdata

import (
	"testing"
	"fmt"
	// "github.com/holmes89/hello-api/translation"
)

func TestRandAdresses(t *testing.T) {

	// Arrange
	anz:=10
	index := "A10"
	
	adrs,err:= LoadRObjects("Adressen",anz,index)

	// Assert
	if err != nil {
		t.Errorf(`random addresses expected - wrong random addresses creation "%s"`, err)
	}
	Report(adrs)
}

// Fügt in Adressen ein Apostroph ein
func TestRandAdressesContent(t *testing.T) {
	// Arrange
	anz:=10
	index := ""
	
	adrs,err:= LoadRObjects("Adressen",anz,index)

	// Assert
	if err != nil {
		t.Errorf(`random addresses expected - wrong random addresses creation "%s"`, err)
	}
	for i, adr := range adrs.Items {
		strasse:=fmt.Sprintf ("%02d %s'%s\n",i,adr.(Adresse).Strasse,"Rest")
		fmt.Printf ("%02d Strasse %s\n",i,strasse)
	
		} 
	// Report(adrs)
}
