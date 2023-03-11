package appdata

import (
	"testing"
	// "github.com/holmes89/hello-api/translation"
)

func TestRandDomainNames(t *testing.T) {

	// Arrange
	var dns RandomObjects
	var res error
	
	// Act
	count := 10
	dns,res = LoadRObjects("DomainNames",count,"")

	// Assert
	if res != nil {
		t.Errorf(`random domain name expected - wrong domain name creation "%s"`, res)
	}
	Report(dns)
}

