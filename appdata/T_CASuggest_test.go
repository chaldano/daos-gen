package appdata

import (
	"testing"
	// "fmt"
	// "github.com/holmes89/hello-api/translation"
)

func TestCASuggest(t *testing.T) {

	// Arrange
	index := "CA"
	count:=120

	co,err := LoadRObjects("CASuggest",count,index)
	
	// Assert
	if err != nil {
		t.Errorf(`random company expected - wrong random company addresses creation "%s"`, err)
	}
	// fmt.Println("Company",co)

	if err == nil{
		Report(co)
	}	
}
