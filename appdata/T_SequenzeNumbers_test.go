package appdata

import (
	"testing"
	// "fmt"
	// "github.com/holmes89/hello-api/translation"
)

func TestRandSequenzes(t *testing.T) {

	// Arrange
	// var seq Zahlensequenzen
	index := "S20"
	count:=20

	// Act
	// aseq := &seq
	// fmt.Println("Durch")
	seq,err2 := LoadRObjects("Zahlensequenz",count,index)
	
	// Assert
	if err2 != nil {
		t.Errorf(`random strings expected - wrong random string creation "%s"`, err2)
	}

	if err2 == nil {
		Report(seq)
	}
}

func TestRandDigits(t *testing.T) {

	// Arrange
	index := ""
	count:=20

	// Act
	seq,err := LoadRObjects("DigitsN",count,index)
	
	// Assert
	if err != nil {
		t.Errorf(`random digest expected - wrong random digest creation "%s"`, err)
	}
	// fmt.Println("DigitsN", seq.Items)
	if err == nil {
		Report(seq)
	}
}