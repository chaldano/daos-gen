package apptests

import (
	"fmt"
	// "regexp"
	"testing"
)
import data "trustkbb.de/tools/dbtools/appdata"


func Partest(s string, a int, i interface{}) (string, error) {
	// fmt.Printf("Anzahl %d für Artikel %s",a,s)
	switch v := i.(type) {
	case data.RandomObjects:
		fmt.Printf("(%v, %T)\n", v, v)
	default:
		fmt.Printf("Anzahl %d für Artikel %s\n", a, s)
	}
	return s, nil
}

func TestParams(t *testing.T) {
	Ware := "Werkzeug"
	Anzahl := 10
	var r data.RandomObjects

	fmt.Println("Ohne Parameter")

	_, err := Partest(Ware, Anzahl, nil)
	if err != nil {
		t.Errorf("Expected Interface value %s\n", err)
	}

	fmt.Println("Mit Parameter")

	_, err = Partest(Ware, Anzahl, r)
	if err != nil {
		t.Errorf("Expected Interface value %s\n", err)
	}
}
