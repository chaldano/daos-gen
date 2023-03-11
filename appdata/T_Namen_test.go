package appdata

import (
	"testing"
	// "fmt"
	// "github.com/holmes89/hello-api/translation"
)

func TestRandNames(t *testing.T) {

	// Arrange
	index := "N10"
	count:=10

	ns,err := LoadRObjects("Namen",count,index)
	

	// Assert
	if err != nil {
		t.Errorf(`random addresses expected - wrong random addresses creation "%s"`, err)
	}
	if err == nil{
		Report(ns)
	}	
}
