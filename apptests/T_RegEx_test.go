package apptests

import (
	"fmt"
	"regexp"
	"testing"
)
import data "trustkbb.de/tools/dbtools/appdata"

func TestDoubleQuoting2(t *testing.T) {
	zeichen := []string{"Hofladen", "Hofl'aden", "Ho'f'laden", "Ho'f'la'den"}
	r := data.ReplaceSingleQuote(zeichen[0])
	if r != zeichen[0] {
		t.Errorf("Expected Error, result must be %s\n", zeichen[0])
	}

	r = data.ReplaceSingleQuote(zeichen[1])
	if r == zeichen[1] {
		t.Errorf("Expected Error, result must be %s\n", "Hofl''aden")
	}
	r = data.ReplaceSingleQuote(zeichen[2])
	fmt.Println("Neue Kette", r)
	if r == zeichen[2] {
		t.Errorf("Expected Error, result must be %s\n", "Ho''f'la''den")
	}
}

func TestMatchString(t *testing.T) {

	zk := []string{"Hallo", "Ha1lo", "Hal2o", "Hall3"}
	for _, s := range zk {
		matched, err := regexp.MatchString(`\d`, s)
		if err != nil {
			t.Errorf("No match found in %s\n", zk)
		}
		fmt.Println("Match", matched)
	}
}
func TestMatchMail(t *testing.T) {

	zk := []string{	"abc@def.com", 
					"abc@def.axc.com", 
					"abc@def.axc.com1", 
					"abc.xyz@def.com", 
					"abc.xyz@de.f.com", 
					"abc.xyz@def.de", 
					"xyz@def.de", 
					"abc.xyz@def.d", 
					"abc.xyz@def.xyz.com", 
					"a.zbc.xyz@def.xyz.com", 
					"def.xyz.com",
				
				}
	for _, s := range zk {
		matched, err := regexp.MatchString(`^[a-zA-Z]+(\.[a-zA-Z]+)?\@\w+\.\w{2,3}$`, s)
		if err != nil {
			t.Errorf("No match found in %s\n", zk)
		}
		if matched != true {
			// fmt.Printf("No-Match: %-5t in %s\n", matched, s)
		} else {
			fmt.Printf("Match   : %-5t in %s\n", matched, s)
		}
	}
}
