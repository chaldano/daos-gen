package appdata

import (
	"testing"
	
)
// import pki "trustkbb.de/tools/dbtools/apppki"


func TestCACertCreation(t *testing.T) {
// Erzeugt eine CA mit X.509 Zertifikat
	// Arrange
	count:=1
	co,err := CreateNObjects("CA",count,"")	
	// Assert
	if err != nil {
		t.Errorf(`CA creation expected - wrong CA creation "%s"`, err)
	}
	if err == nil{
		Report(co)
	}	
}

func TestUserCertCreation(t *testing.T) {

	// Arrange
	count:=1
	co,err := CreateNObjects("CA",count,"")	
	// Assert
	if err != nil {
		t.Errorf(`CA creation expected - wrong CA creation "%s"`, err)
	}
	if err == nil{
		Report(co)
	}
	
	ca:=co.Items[0].(CA)
	
	cop,err := CreateNObjects("PersonsNat",count,"")	
	// Assert
	if err != nil {
		t.Errorf(`Person creation expected - wrong Person creation "%s"`, err)
	}
	// Für jede CA ein CA-Zertifikat erzeugen
	for i, p :=range cop.Items{
	    person:= p.(PersonNat)
		err = ca.issueCert(&person)
		if err != nil {
			t.Errorf(`Person certificate creation expected - wrong Person certifictae creation "%s"`, err)
		}
		cop.Items[i] = person
	}
	
	if err == nil{
		Report(cop)
	}	
}