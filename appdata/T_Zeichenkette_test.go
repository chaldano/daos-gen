package appdata

import (
	"testing"
	// "github.com/holmes89/hello-api/translation"
)

func TestRandStrings(t *testing.T) {

	// Arrange
	index := "Z10"
	count:=10


	zks,res := LoadRObjects("Zeichenketten",count,index)

	// Assert
	if res != nil {
		t.Errorf(`random strings expected - wrong random string creation "%s"`, res)
	}

	if res == nil {
		Report(zks)
	}


}
