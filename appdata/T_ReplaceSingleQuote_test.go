package appdata

import (
	"testing"
	"fmt"
)

func TestDoubleQuoting (t *testing.T) {	
zeichen  := []string {"Hofladen","Hofl'aden","Ho'f'laden","Ho'f'la'den"}
r := ReplaceSingleQuote(zeichen[0])
if r != zeichen[0] {
	t.Errorf("Expected Error, result must be %s\n", zeichen[0])
	}

r = ReplaceSingleQuote(zeichen[1])
if r ==  zeichen[1] {
	t.Errorf("Expected Error, result must be %s\n", "Hofl''aden")
	}
r = ReplaceSingleQuote(zeichen[2])
fmt.Println("Neue Kette",r)
	if r ==  zeichen[2] {
		t.Errorf("Expected Error, result must be %s\n", "Ho''f'la''den")
		}	
}