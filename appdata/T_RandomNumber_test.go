package appdata

import (
	"testing"
	"fmt"
	// "github.com/holmes89/hello-api/translation"
)

func TestRandDigitFrom(t *testing.T) {
	// Arrange
	count:=5

	// Act
	number,err := RandDigitFrom(count)
	
	// Assert
	if err != nil {
		t.Errorf(`random digest expected - wrong random digest creation "%s"`, err)
	}
	fmt.Printf("Random Digit from %d : %d\n", count,number)
}