package appdata

import (
	"testing"
	// "fmt"
)

func TestSocialString(t *testing.T) {
	
	count:=10
	sns,err := CreateNObjects("RegisterNumbers",count,nil)
	
	// Assert
	if err != nil {
		t.Errorf(`random socialNumbersStrings expected - wrong random socialNumbersString creation "%s"`, err)
	
	}
	 if err == nil {
	 Report(sns)
	}

}

