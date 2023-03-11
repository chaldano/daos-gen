package appdata

import (
	"testing"
	// "fmt"
	// "github.com/holmes89/hello-api/translation"
)

func TestRegisterNumber(t *testing.T) {
	index := "SN01"
	count:=1

	sns,err2 := LoadRObjects("RegisterNumbers",count,index)

	// Assert
	if err2 != nil {
		t.Errorf(`random RegisterNumbers expected - wrong random RegisterNumbers creation "%s"`, err2)
	
	}
	if err2 == nil {
	Report(sns)
	}

}
	
