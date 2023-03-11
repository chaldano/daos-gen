package appdata

import (
	"testing"
)

func TestRandUuids(t *testing.T) {

	// Arrange
	// Act
	count := 10
	index := ""

	uuids,res := LoadRObjects("UUID",count, index)
	// Assert
	if res != nil {
	t.Errorf(`random uuids expected - wrong random uuid creation "%s"`, res)
	}

Report(uuids)

}
